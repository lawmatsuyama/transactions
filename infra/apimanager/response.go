package apimanager

import (
	"github.com/lawmatsuyama/transactions/domain"
)

type Transaction struct {
	Description string
	Amount      float64
	Operation   string
}

type TransactionSaveResponse struct {
	Transaction Transaction
	Errors      []string
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
	Error  string `json:"error"`
	Result any    `json:"result"`
}
