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
	_, err := CreateQueue(domain.QueueTransactionSaved, true, arg)
	if err != nil {
		return err
	}

	err = CreateExchange(domain.ExchangeTransactionSaved, "fanout", nil)
	if err != nil {
		return err
	}

	return BindQueueExchange(domain.QueueTransactionSaved, domain.ExchangeTransactionSaved, "transaction-service")
}
