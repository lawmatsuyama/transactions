package apimanager

import (
	"time"

	"github.com/lawmatsuyama/transactions/domain"
)

// TransactionSaveRequest represents a request of Save transaction operation
type TransactionSaveRequest struct {
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
	Operation   string  `json:"operation" bson:"operation"`
}

// TransactionSaveRequest represents a request of Save transaction operation represents a request of Save transaction operation
type TransactionsSaveRequest struct {
	UserID        string                   `json:"user_id" bson:"user_id"`
	OriginChannel string                   `json:"origin_channel" bson:"origin_channel"`
	Transactions  []TransactionSaveRequest `json:"transactions" bson:"transactions"`
}

// ToTransactions return a new domain Transactions from TransactionsSaveRequest
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
	Page     int64  `json:"page"`
	NextPage *int64 `json:"next_page,omitempty"`
}

// TransactionsGetRequest represents a request of Get transactions operation
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

// ToTransactionsFilter return a new domain TransactionsFilter from TransactionsGetRequest
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
