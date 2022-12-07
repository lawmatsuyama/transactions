package domain

type OperationType string

const (
	DebitOperation  OperationType = "debit"
	CreditOperation OperationType = "credit"
)
