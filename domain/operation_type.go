package domain

// OperationType represents an enum of operation type
type OperationType string

const (
	DebitOperation  OperationType = "debit"
	CreditOperation OperationType = "credit"
)

// IsValid check if operation type is valid
func (oper OperationType) IsValid() bool {
	switch oper {
	case DebitOperation, CreditOperation:
		return true
	default:
		return false
	}
}
