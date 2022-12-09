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

type Paging struct {
	Page int64
}

type TransactionsGetRequest struct {
	ID            string    `json:"_id"`
	Description   string    `json:"description"`
	UserID        string    `json:"user_id"`
	Origin        string    `json:"origin"`
	OperationType string    `json:"operation_type"`
	AmountGreater float64   `json:"amount_greater"`
	AmountLess    float64   `json:"amount_less"`
	DateFrom      time.Time `json:"date_from"`
	DateTo        time.Time `json:"date_to"`
	Paging        *Paging   `json:"paging,omitempty"`
}

func (req TransactionsGetRequest) ToTransactionsFilter() domain.TransactionFilter {
	filter := domain.TransactionFilter{
		ID:            req.ID,
		Description:   req.Description,
		UserID:        req.UserID,
		Origin:        domain.OriginChannel(req.Origin),
		OperationType: domain.OperationType(req.OperationType),
		AmountGreater: req.AmountGreater,
		AmountLess:    req.AmountLess,
		DateFrom:      req.DateFrom,
		DateTo:        req.DateTo,
	}

	if req.Paging != nil {
		filter.Paging = &domain.Paging{
			Page: req.Paging.Page,
		}
	}

	return filter
}
