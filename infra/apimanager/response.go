package apimanager

import (
	"github.com/lawmatsuyama/transactions/domain"
)

type Transaction struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Operation   string  `json:"operation"`
}

type TransactionSaveResponse struct {
	Transaction Transaction `json:"transaction"`
	Errors      []string    `json:"errors"`
}

func FromTransactionValidateResult(trsResult []domain.TransactionValidateResult) []TransactionSaveResponse {
	trs := make([]TransactionSaveResponse, len(trsResult))
	for i, trResult := range trsResult {
		trs[i] = TransactionSaveResponse{
			Transaction: Transaction{
				Description: trResult.Transaction.Description,
				Amount:      trResult.Transaction.Amount,
				Operation:   string(trResult.Transaction.OperationType),
			},
			Errors: trResult.Errors,
		}
	}
	return trs
}

type GenericResponse struct {
	Error  string `json:"error,omitempty"`
	Result any    `json:"result"`
}
