package domain

// TransactionSaveResult represents an result of operation Save
type TransactionSaveResult struct {
	Transaction *Transaction `json:"transaction"`
	Errors      []string     `json:"errors"`
}
