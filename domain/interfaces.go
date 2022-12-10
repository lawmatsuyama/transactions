package domain

import (
	"context"
)

// UserRepository represents user repository
type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
}

// TransactionRepository represents transaction repository
type TransactionRepository interface {
	Save(ctx context.Context, transactions Transactions) error
	Get(ctx context.Context, filterTrs TransactionFilter) (TransactionsPaging, error)
}

// SessionControlRepository represents session control repository
type SessionControlRepository interface {
	WithSession(ctx context.Context, f FuncDBSession) error
}

// MessagePublisher represents message publisher repository
type MessagePublisher interface {
	Publish(ctx context.Context, excName, routingKey string, obj any, priority uint8) error
}

// TransactionUseCase represents transaction use case
type TransactionUseCase interface {
	Save(ctx context.Context, userID string, transactions Transactions) ([]TransactionSaveResult, error)
	Get(ctx context.Context, filterTrs TransactionFilter) (TransactionsPaging, error)
}

// UUIDGenerator represents an UUID generator
type UUIDGenerator interface {
	Generate() string
}

// IDSetter represents an ID setter
type IDSetter interface {
	SetID()
}
