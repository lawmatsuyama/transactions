package apimanager

import (
	"github.com/lawmatsuyama/transactions/domain"
)

type GenericResponse struct {
	Error  string `json:"error,omitempty"`
	Result any    `json:"result"`
}

type Transaction struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Operation   string  `json:"operation"`
	Origin      string  `json:"origin"`
	CreatedAt   string  `json:"created_at"`
}

type TransactionSaveResponse struct {
	Transaction Transaction `json:"transaction"`
	Errors      []string    `json:"errors"`
}

func FromTransactionSaveResult(trsResult []domain.TransactionSaveResult) []TransactionSaveResponse {
	trs := make([]TransactionSaveResponse, len(trsResult))
	for i, trResult := range trsResult {
		trs[i] = TransactionSaveResponse{
			Transaction: Transaction{
				ID:          trResult.Transaction.ID,
				UserID:      trResult.Transaction.UserID,
				Origin:      string(trResult.Transaction.Origin),
				Description: trResult.Transaction.Description,
				Amount:      trResult.Transaction.Amount,
				Operation:   string(trResult.Transaction.OperationType),
				CreatedAt:   trResult.Transaction.CreatedAt.Format("2006-01-02T15-04-05"),
			},
			Errors: trResult.Errors,
		}
	}
	return trs
}

type TransactionsGetResponse struct {
	Transactions []Transaction `json:"transactions"`
	Paging       *Paging       `json:"paging"`
}

func FromTransactionPaging(trsPag domain.TransactionsPaging) TransactionsGetResponse {
	trs := make([]Transaction, len(trsPag.Transactions))
	for i, tr := range trsPag.Transactions {
		trs[i] = Transaction{
			ID:          tr.ID,
			UserID:      tr.UserID,
			Description: tr.Description,
			Amount:      tr.Amount,
			Operation:   string(tr.OperationType),
			Origin:      string(tr.Origin),
			CreatedAt:   tr.CreatedAt.Format("2006-01-02T15-04-05"),
		}
	}

	trsGetResp := TransactionsGetResponse{Transactions: trs}
	if trsPag.Paging != nil {
		trsGetResp.Paging = &Paging{
			Page: trsPag.Paging.Page,
		}
	}

	return trsGetResp
}
