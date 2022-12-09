package messagebroker

import (
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/streadway/amqp"
)

type Setuper struct{}

func NewSetuper() Setuper {
	return Setuper{}
}

func (setup Setuper) Setup() error {
	arg := amqp.Table{"x-max-priority": 9}
	_, err := CreateQueue(domain.QueueSaved, true, arg)
	if err != nil {
		return err
	}

	err = CreateExchange(domain.ExchangeTransaction, "fanout", nil)
	if err != nil {
		return err
	}

	return BindQueueExchange(domain.QueueSaved, domain.ExchangeTransaction, "transaction-service")
}
