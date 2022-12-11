package messagebroker

// BrokerSetuper represents a setuper to config message broker
type BrokerSetuper interface {
	Setup() error
}
