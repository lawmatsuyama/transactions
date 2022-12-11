package domain

// BrokerSetup represents a function to setup message broker
type BrokerSetup func() error

// Setup execute setup function to configure message broker
func (s BrokerSetup) Setup() error {
	return s()
}

const (
	// ExchangeTransactionSaved is a exchange to notify transactions saved
	ExchangeTransactionSaved = "EXC.TRANSACTION.SAVED"
	// QueueTransactionSaved is one of the queues that receive transactions saved
	QueueTransactionSaved = "QUEUE.TRANSACTION.SAVED"
)
