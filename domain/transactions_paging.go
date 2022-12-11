package domain

var (
	LimitTransactionsByPage int64 = 20
)

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

// SetNextPage set next page
func (p *Paging) SetNextPage(numTransactions int) {
	if p == nil {
		return
	}

	if int64(numTransactions) >= LimitTransactionsByPage {
		nextPage := p.Page + 1
		p.NextPage = &nextPage
	}
}

func (p *Paging) TransactionsSkip() int64 {
	page := p.CurrentPage()
	return (page - 1) * LimitTransactionsByPage
}

func (p *Paging) LimitTransactionsByPage() int64 {
	return LimitTransactionsByPage
}

type TransactionsPaging struct {
	Transactions Transactions `json:"transactions"`
	Paging       *Paging      `json:"paging,omitempty"`
}

func (trs TransactionsPaging) IsValid() error {
	if len(trs.Transactions) == 0 {
		return ErrTransactionsNotFound
	}

	return nil
}

func (trs *TransactionsPaging) SetNextPaging() {
	if trs == nil {
		return
	}

	if trs.Paging == nil {
		trs.Paging = &Paging{Page: 1}
	}

	trs.Paging.SetNextPage(len(trs.Transactions))
}
