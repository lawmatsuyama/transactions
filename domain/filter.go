package domain

import "time"

type TransactionFilter struct {
	ID            string        `json:"_id"`
	Description   string        `json:"description"`
	UserID        string        `json:"user_id"`
	Origin        OriginChannel `json:"origin"`
	OperationType OperationType `json:"operation_type"`
	AmountGreater float64       `json:"amount_greater"`
	AmountLess    float64       `json:"amount_less"`
	DateFrom      time.Time     `json:"date_from"`
	DateTo        time.Time     `json:"date_to"`
	Paging        *Paging       `json:"paging,omitempty"`
}

func (tr TransactionFilter) Validate() error {
	if tr.OperationType != "" && !tr.OperationType.IsValid() {
		return ErrInvalidOperationType
	}

	if tr.Origin != "" && !tr.Origin.IsValid() {
		return ErrInvalidOriginChannel
	}

	return nil
}

func (tr TransactionFilter) Page() int64 {
	if tr.Paging == nil {
		return 0
	}

	return tr.Paging.Page
}
