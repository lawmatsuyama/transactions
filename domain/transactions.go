package domain

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	LimitTransactionsToSave int = 20
)

// Transactions represents a list of transaction
type Transactions []*Transaction

// ValidateTransactionsToSave check if all transactions are valid
func (trs Transactions) ValidateTransactionsToSave() ([]TransactionSaveResult, error) {
	chResult := make(chan TransactionSaveResult, len(trs))
	trsResult := make([]TransactionSaveResult, 0, len(trs))

	if len(trs) == 0 {
		log.WithError(ErrInvalidTransaction).Error("There is no transactions to process")
		return nil, ErrInvalidTransaction
	}

	if len(trs) > LimitTransactionsToSave {
		log.WithField("limit_transactions", LimitTransactionsToSave).WithError(ErrInvalidTransaction).Error("Too many transactions to save")
		return nil, ErrTooManyTransaction
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
