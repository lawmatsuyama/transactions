package apimanager

import (
	"time"

	"github.com/lawmatsuyama/transactions/domain"
)

type TransactionSaveRequest struct {
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
	Operation   string  `json:"operation" bson:"operation"`
}

type TransactionsSaveRequest struct {
	UserID        string                   `json:"user_id" bson:"user_id"`
	OriginChannel string                   `json:"origin_channel" bson:"origin_channel"`
	Transactions  []TransactionSaveRequest `json:"transactions" bson:"transactions"`
}

func (trsReq TransactionsSaveRequest) ToTransactions(now time.Time) domain.Transactions {
	trs := make([]*domain.Transaction, len(trsReq.Transactions))
	for i, trReq := range trsReq.Transactions {
		trs[i] = &domain.Transaction{
			Description:   trReq.Description,
			UserID:        trsReq.UserID,
			Origin:        domain.OriginChannel(trsReq.OriginChannel),
			Amount:        trReq.Amount,
			OperationType: domain.OperationType(trReq.Operation),
			CreatedAt:     now,
		}
	}
	return trs
}
