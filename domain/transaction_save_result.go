package domain

type TransactionSaveResult struct {
	Transaction *Transaction `json:"transaction"`
	Errors      []string     `json:"errors"`
}
