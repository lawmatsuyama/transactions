package messagebroker

import (
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/streadway/amqp"
)

// Setuper implements BrokerSetuper interface
type Setuper struct{}

// NewSetuper returns a new Setuper
func NewSetuper() Setuper {
	return Setuper{}
}

// Setup configure and bind queue and exchange
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
