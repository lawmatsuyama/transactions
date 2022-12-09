package domain

type BrokerSetup func() error

func (s BrokerSetup) Setup() error {
	return s()
}

const (
	ExchangeTransaction = "EXC.TRANSACTION.SAVED"
	QueueSaved          = "QUEUE.TRANSACTION.SAVED"
)
