package domain

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
}

type TransactionRepository interface {
	Save(ctx context.Context, transactions Transactions) error
	Get(ctx context.Context, filterTrs TransactionFilter) (TransactionsPaging, error)
}

type SessionControlRepository interface {
	WithSession(ctx context.Context, f FuncDBSession) error
}

type MessagePublisher interface {
	Publish(ctx context.Context, excName, routingKey string, obj any, priority uint8) error
}

type TransactionUseCase interface {
	Save(ctx context.Context, userID string, transactions Transactions) ([]TransactionSaveResult, error)
	Get(ctx context.Context, filterTrs TransactionFilter) (TransactionsPaging, error)
}

type UUIDGenerator interface {
	Generate() string
}

type IDSetter interface {
	SetID()
}
