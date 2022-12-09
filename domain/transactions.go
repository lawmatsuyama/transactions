package domain

import (
	"sync"
)

type Transactions []*Transaction

func (trs Transactions) ValidateTransactions() ([]TransactionValidateResult, error) {
	chResult := make(chan TransactionValidateResult, len(trs))
	trsResult := make([]TransactionValidateResult, 0, len(trs))
	go func() {
		trs.validateTransactionsConcurrently(chResult)
	}()

	for result := range chResult {
		if len(result.Errors) > 0 {
			trsResult = append(trsResult, result)
		}
	}

	if len(trsResult) > 0 {
		return trsResult, ErrInvalidTransaction
	}

	return nil, nil
}

func (trs Transactions) validateTransactionsConcurrently(chResult chan<- TransactionValidateResult) {
	wg := &sync.WaitGroup{}
	for _, tr := range trs {
		wg.Add(1)
		go func(tr *Transaction) {
			defer wg.Done()
			if listErr := tr.IsValid(); len(listErr) > 0 {
				chResult <- TransactionValidateResult{Transaction: tr, Errors: listErr}
			}
		}(tr)
	}
	wg.Wait()
	close(chResult)
}
