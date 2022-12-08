package domain

type TransactionValidateResult struct {
	Transaction Transaction
	Error       error
}
