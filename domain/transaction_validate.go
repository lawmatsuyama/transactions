package domain

type TransactionValidateResult struct {
	Transaction *Transaction `json:"transaction"`
	Errors      []string     `json:"errors"`
}
