package domain

var (
	// LimitTransactionsByPage is a global variable to limit the transactions of listing page
	LimitTransactionsByPage int64 = 20
)

// Paging represents paging struct
type Paging struct {
	Page     int64  `json:"page"`
	NextPage *int64 `json:"next_page,omitempty"`
}

// CurrentPage returns the current page number
func (p *Paging) CurrentPage() int64 {
	if p == nil || p.Page <= 0 {
		return 1
	}

	return p.Page
}

// SetNextPage set next page according to num transactions
func (p *Paging) SetNextPage(numTransactions int) {
	if p == nil {
		return
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	if int64(numTransactions) >= LimitTransactionsByPage {
		nextPage := p.Page + 1
		p.NextPage = &nextPage
	}
}

// TransactionsSkip returns the number of transaction to skip to view the page
func (p *Paging) TransactionsSkip() int64 {
	page := p.CurrentPage()
	return (page - 1) * LimitTransactionsByPage
}

// LimitTransactionsByPage returns the limit number of transactions by page
func (p *Paging) LimitTransactionsByPage() int64 {
	return LimitTransactionsByPage
}

// TransactionsPaging represents transactions by page
type TransactionsPaging struct {
	Transactions Transactions `json:"transactions"`
	Paging       *Paging      `json:"paging,omitempty"`
}

// IsValid check if TransactionsPaging is valid
func (trs TransactionsPaging) IsValid() error {
	if len(trs.Transactions) == 0 {
		return ErrTransactionsNotFound
	}

	return nil
}

// SetNextPaging set the next paging according to number of transactions
func (trs *TransactionsPaging) SetNextPaging() {
	if trs == nil {
		return
	}

	if trs.Paging == nil {
		trs.Paging = &Paging{}
	}

	trs.Paging.SetNextPage(len(trs.Transactions))
}
