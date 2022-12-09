package domain

type Paging struct {
	Page int64 `json:"page"`
}

type TransactionsPaging struct {
	Transactions Transactions `json:"transactions"`
	Paging       *Paging      `json:"paging,omitempty"`
}
