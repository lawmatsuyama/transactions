package usecases_test

import (
	"context"

	"github.com/lawmatsuyama/transactions/domain"
)

var (
	SaveMock        func(ctx context.Context, transactions []*domain.Transaction) error
	GetByIDMock     func(ctx context.Context, id string) (domain.User, error)
	PublishMock     func(ctx context.Context, excName, routingKey string, obj any, priority uint8) error
	WithSessionMock func(ctx context.Context, f domain.FuncDBSession) error
)

type mock struct{}

func (m mock) Save(ctx context.Context, transactions []*domain.Transaction) error {
	return SaveMock(ctx, transactions)
}

func (m mock) GetByID(ctx context.Context, id string) (domain.User, error) {
	return GetByIDMock(ctx, id)
}

func (m mock) Publish(ctx context.Context, excName, routingKey string, obj any, priority uint8) error {
	return PublishMock(ctx, excName, routingKey, obj, priority)
}

func (m mock) WithSession(ctx context.Context, f domain.FuncDBSession) error {
	return WithSessionMock(ctx, f)
}
