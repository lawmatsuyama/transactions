package messagebroker

import "context"

// MessagePublisher it is rabbitmq implementation of domain.MessagePublisher interface
type MessagePublisher struct{}

// NewMessagePublisher returns a new message publisher
func NewMessagePublisher() MessagePublisher {
	return MessagePublisher{}
}

// Publish publish message given an exchange and routingKey
func (pub MessagePublisher) Publish(ctx context.Context, excName, routingKey string, obj any, priority uint8) error {
	return Publish(ctx, excName, routingKey, obj, priority)
}
