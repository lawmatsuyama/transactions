package messagebroker

import "context"

type MessagePublisher struct{}

func NewMessagePublisher() MessagePublisher {
	return MessagePublisher{}
}

func (pub MessagePublisher) Publish(ctx context.Context, excName, routingKey string, obj interface{}, priority uint8) error {
	return Publish(ctx, excName, routingKey, obj, priority)
}
