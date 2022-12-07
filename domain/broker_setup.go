package domain

type BrokerSetup func() error

func (s BrokerSetup) Setup() error {
	return s()
}
