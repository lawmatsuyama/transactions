package messagebroker

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/lawmatsuyama/transactions/domain"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type messageBroker struct {
	connection         *amqp.Connection
	publisherChannel   *amqp.Channel
	consumerChannel    *amqp.Channel
	notifyCloseChannel chan *amqp.Error
}

var (
	broker              *messageBroker
	delayToConnectAgain = time.Second * 5
)

var ErrMessageBrokerClosed = errors.New("message broker is closed")

func newMessageBroker(conn *amqp.Connection, chPub, chCons *amqp.Channel, chNotify chan *amqp.Error) *messageBroker {
	return &messageBroker{
		connection:         conn,
		publisherChannel:   chPub,
		consumerChannel:    chCons,
		notifyCloseChannel: chNotify,
	}
}

// Start start connection and setup queues/exchanges according to the given setuper.
func Start(ctx context.Context, setuper BrokerSetuper) {
	url := os.Getenv("MESSAGE_BROKER_URL")
	if url == "" {
		panic("message broker url is empty")
	}

	go func() {
		for {
			err := Connect(url)
			if err != nil {
				defaultSleep()
				continue
			}

			if setuper != nil {
				err = setuper.Setup()
				if err != nil {
					defaultSleep()
					continue
				}
			}

			select {
			case <-broker.notifyCloseChannel:
				defaultSleep()
			case <-ctx.Done():
				log.Info("message broker stopped")
				return
			}
		}
	}()
}

// Shutdown finishes connection with message broker
func Shutdown() {
	if broker == nil || broker.connection == nil {
		return
	}

	err := broker.connection.Close()
	if err != nil {
		log.Warn("Failed to close rabbitmq connection")
	}
}

// Connect connect to rabbitmq and open channels
func Connect(url string) error {
	conn, err := connect(url)
	if err != nil {
		return err
	}

	notifyClose := conn.NotifyClose(make(chan *amqp.Error))

	chPub, chCons, err := openChannels(conn)
	if err != nil {
		return err
	}

	broker = newMessageBroker(conn, chPub, chCons, notifyClose)
	return nil
}

// CreateQueue create queue on rabbitmq. If the queue does not already exist, the server will create it
func CreateQueue(queueName string, durable bool, args amqp.Table) (amqp.Queue, error) {
	q, err := broker.consumerChannel.QueueDeclare(
		queueName,
		durable,
		false,
		!durable,
		false,
		args,
	)
	return q, err
}

// CreateExchange creates exchange on rabbitMQ. If the exchange does not already exist, the server will create it
func CreateExchange(excName string, excType string, args amqp.Table) error {
	err := broker.consumerChannel.ExchangeDeclare(
		excName,
		excType,
		true,
		false,
		false,
		false,
		args,
	)

	return err
}

// BindQueueExchange bind an exchange to the queue.
func BindQueueExchange(queueName, exchangeName, routingKey string) error {
	return broker.consumerChannel.QueueBind(
		queueName,    //name of the queue
		routingKey,   //routing key
		exchangeName, //name of the exchange
		true,         //no wait
		nil,          //arguments
	)
}

// Consume create consumer to queue
func Consume(ctx context.Context, queueName, consumer string, f func(amqp.Delivery)) {
	go func() {
		consume(ctx, queueName, consumer, f)
	}()
}

// Publish publish message by given exchange and routingKey
func Publish(ctx context.Context, excName, routingKey string, obj any, priority uint8) error {
	retry := 0
	for {
		if broker == nil {
			retry++
			if retry > 4 {
				return ErrMessageBrokerClosed
			}
			defaultSleep()
			continue
		}

		select {
		case <-broker.notifyCloseChannel:
			retry++
			if retry > 4 {
				return ErrMessageBrokerClosed
			}
			defaultSleep()
			continue
		default:
		}

		break
	}

	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return broker.publisherChannel.Publish(
		excName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/javascript",
			Body:        body,
			Priority:    priority,
		},
	)
}

func consume(ctx context.Context, queueName, consumer string, f func(amqp.Delivery)) {
	msgs, err := broker.consumerChannel.Consume(
		queueName,
		consumer,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.WithFields(log.Fields{"consumer": consumer, "queue": queueName}).
			WithError(err).Error("Failed to consume message")
		return
	}

	for {
		select {
		case msg := <-msgs:
			domain.AddTaskCount()
			go func() {
				defer domain.DoneTask()
				f(msg)
			}()
		case <-broker.notifyCloseChannel:
			return
		case <-ctx.Done():
			cancelConsumer(consumer)
			log.WithField("consumer", consumer).Info("Message broker consumer stopped")
			return
		}
	}
}

func connect(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.WithField("url", url).WithError(err).Error("Could not connect to rabbitmq")
		return nil, err
	}

	log.WithField("url", url).Info("Rabbitmq connected")
	return conn, err
}

func openChannels(conn *amqp.Connection) (chPub, chCons *amqp.Channel, err error) {
	chPub, err = conn.Channel()
	if err != nil {
		log.WithError(err).Error("Could not connect to publisher channel")
		return
	}

	chCons, err = conn.Channel()
	if err != nil {
		log.WithError(err).Error("Could not connect to consumer channel")
		return
	}

	log.Info("channels opened")
	return
}

func defaultSleep() {
	time.Sleep(delayToConnectAgain)
}

func cancelConsumer(consumer string) {
	err := broker.consumerChannel.Cancel(consumer, false)
	if err != nil {
		log.WithField("consumer", consumer).WithError(err).Error("Failed to cancel consumer")
	}
}
