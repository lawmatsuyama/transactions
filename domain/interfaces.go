package domain

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
}

type TransactionRepository interface {
	Save(ctx context.Context, transactions []*Transaction) error
	// GetID(ctx context.Context, id string) (Transaction, error)
	// GetByUserID(ctx context.Context, userID string, page int) ([]Transaction, error)
	// GetByUserIDAndID(ctx context.Context, id, userID string) (Transaction, error)
	// GetByUserIDAndFromDate(ctx context.Context, userID string, date time.Time, page int) ([]Transaction, error)
	// GetFromDate(ctx context.Context, date time.Time, page int) ([]Transaction, error)
}

type MessagePublisher interface {
	Publish(ctx context.Context, excName, routingKey string, obj interface{}, priority uint8) error
}

type TransactionUseCase interface {
	Save(ctx context.Context, userID string, transactions Transactions) ([]TransactionValidateResult, error)
}

type UUIDGenerator interface {
	Generate() string
}

type IDSetter interface {
	SetID()
}
