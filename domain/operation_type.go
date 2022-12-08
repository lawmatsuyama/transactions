package domain

type OperationType string

const (
	DebitOperation  OperationType = "debit"
	CreditOperation OperationType = "credit"
)

func (oper OperationType) IsValid() bool {
	switch oper {
	case DebitOperation, CreditOperation:
		return true
	default:
		return false
	}
}
