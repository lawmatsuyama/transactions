package domain

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// Transactions represents a list of transaction
type Transactions []*Transaction

// ValidateTransactions check if all transactions are valid
func (trs Transactions) ValidateTransactions() ([]TransactionSaveResult, error) {
	chResult := make(chan TransactionSaveResult, len(trs))
	trsResult := make([]TransactionSaveResult, 0, len(trs))

	if len(trs) == 0 {
		logrus.WithError(ErrInvalidTransaction).Error("There is no transactions to process")
		return nil, ErrInvalidTransaction
	}

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

func (trs Transactions) validateTransactionsConcurrently(chResult chan<- TransactionSaveResult) {
	wg := &sync.WaitGroup{}
	for _, tr := range trs {
		wg.Add(1)
		go func(tr *Transaction) {
			defer wg.Done()
			if listErr := tr.IsValid(); len(listErr) > 0 {
				chResult <- TransactionSaveResult{Transaction: tr, Errors: listErr}
			}
		}(tr)
	}
	wg.Wait()
	close(chResult)
}
