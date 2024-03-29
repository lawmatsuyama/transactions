package domain

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// Transaction represents a transaction of an user
type Transaction struct {
	ID            string        `json:"_id" bson:"_id,omitempty"`
	Description   string        `json:"description" bson:"description"`
	UserID        string        `json:"user_id" bson:"user_id"`
	Origin        OriginChannel `json:"origin" bson:"origin"`
	Amount        float64       `json:"amount" bson:"amount"`
	OperationType OperationType `json:"operation_type" bson:"operation_type"`
	CreatedAt     time.Time     `json:"created_at" bson:"created_at"`
}

// IsAmountZero return true if transaction amount is zero
func (tr Transaction) IsAmountZero() bool {
	return tr.Amount == 0
}

// IsValid check if transaction is valid
func (tr Transaction) IsValid() []string {
	listErr := []string{}
	if tr.IsAmountZero() {
		listErr = append(listErr, ErrTransactionZeroAmount.Error())
	}

	if !tr.OperationType.IsValid() {
		listErr = append(listErr, ErrInvalidOperationType.Error())
	}

	if !tr.Origin.IsValid() {
		listErr = append(listErr, ErrInvalidOriginChannel.Error())
	}

	return listErr
}

// SetID generate and set ID for transaction
func (tr *Transaction) SetID() {
	if tr == nil {
		log.Error("cannot set ID for transaction because transaction is nil")
		return
	}

	if tr.Description != "" {
		tr.ID = fmt.Sprintf("%s-%s-%s-%s-%s-%s", tr.UserID, tr.OperationType, tr.Origin, TimeSaoPaulo(tr.CreatedAt).Format("2006-01-02"), tr.Description, UUID.Generate())
		return
	}

	tr.ID = fmt.Sprintf("%s-%s-%s-%s-%s", tr.UserID, tr.OperationType, tr.Origin, TimeSaoPaulo(tr.CreatedAt).Format("2006-01-02"), UUID.Generate())
}
