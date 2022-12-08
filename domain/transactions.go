package domain

import (
	"sync"
)

type Transactions []Transaction

func (trs Transactions) ValidateTransactions() ([]TransactionValidateResult, error) {
	chResult := make(chan TransactionValidateResult, len(trs))
	trsResult := make([]TransactionValidateResult, 0, len(trs))
	go func() {
		trs.validateTransactionsConcurrently(chResult)
	}()

	for result := range chResult {
		if result.Error != nil {
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
		go func(tr Transaction) {
			defer wg.Done()
			if err := tr.IsValid(); err != nil {
				chResult <- TransactionValidateResult{Transaction: tr, Error: err}
			}
		}(tr)
	}
	wg.Wait()
	close(chResult)
}
