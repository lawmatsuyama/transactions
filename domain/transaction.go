package domain

import "time"

type Transaction struct {
	ID            string        `json:"_id" bson:"_id"`
	Description   string        `json:"description" bson:"description"`
	UserID        string        `json:"user_id" bson:"user_id"`
	Origin        OriginChannel `json:"origin" bson:"origin"`
	Amount        float64       `json:"amount" bson:"amount"`
	OperationType OperationType `json:"operation_type" bson:"operation_type"`
	CreatedAt     time.Time     `json:"created_at" bson:"created_at"`
}

func (tr Transaction) IsAmountZero() bool {
	return tr.Amount == 0
}

func (tr Transaction) IsValid() error {
	if !tr.IsAmountZero() {
		return ErrTransactionZeroAmount
	}

	if !tr.OperationType.IsValid() {
		return ErrInvalidOperationType
	}

	if !tr.Origin.IsValid() {
		return ErrInvalidOriginChannel
	}

	return nil
}
