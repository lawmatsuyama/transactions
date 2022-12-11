package domain_test

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateTransactions(t *testing.T) {
	testCases := []struct {
		Name                               string
		TransactionsFile                   string
		ExepectedTransactionSaveResultFile string
		LimitTransactionsToSave            int
		ExpectedError                      error
	}{
		{
			Name:                               "01_should_validate_transactions_return_error_nil",
			TransactionsFile:                   "./testdata/_transactions/01_should_validate_transactions_return_error_nil/transactions.json",
			ExepectedTransactionSaveResultFile: "./testdata/_transactions/01_should_validate_transactions_return_error_nil/exp_transactions_result.json",
			LimitTransactionsToSave:            30,
			ExpectedError:                      nil,
		},
		{
			Name:                               "02_should_validate_transactions_return_invalid_operation",
			TransactionsFile:                   "./testdata/_transactions/02_should_validate_transactions_return_invalid_operation/transactions.json",
			ExepectedTransactionSaveResultFile: "./testdata/_transactions/02_should_validate_transactions_return_invalid_operation/exp_transactions_result.json",
			LimitTransactionsToSave:            30,
			ExpectedError:                      domain.ErrInvalidTransaction,
		},
		{
			Name:                               "03_should_validate_transactions_return_invalid_origin_and_operation",
			TransactionsFile:                   "./testdata/_transactions/03_should_validate_transactions_return_invalid_origin_and_operation/transactions.json",
			ExepectedTransactionSaveResultFile: "./testdata/_transactions/03_should_validate_transactions_return_invalid_origin_and_operation/exp_transactions_result.json",
			LimitTransactionsToSave:            30,
			ExpectedError:                      domain.ErrInvalidTransaction,
		},
		{
			Name:                               "05_should_validate_transactions_return_too_many_transactions",
			TransactionsFile:                   "./testdata/_transactions/04_no_transactions_should_validate_transactions_return_invalid_transaction/transactions.json",
			ExepectedTransactionSaveResultFile: "./testdata/_transactions/04_no_transactions_should_validate_transactions_return_invalid_transaction/exp_transactions_result.json",
			LimitTransactionsToSave:            30,
			ExpectedError:                      domain.ErrInvalidTransaction,
		},
		{
			Name:                               "05_should_validate_transactions_return_too_many_transactions",
			TransactionsFile:                   "./testdata/_transactions/05_should_validate_transactions_return_too_many_transactions/transactions.json",
			ExepectedTransactionSaveResultFile: "./testdata/_transactions/05_should_validate_transactions_return_too_many_transactions/exp_transactions_result.json",
			LimitTransactionsToSave:            4,
			ExpectedError:                      domain.ErrTooManyTransaction,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testValidateTransactions(t, tc.Name, tc.TransactionsFile, tc.ExepectedTransactionSaveResultFile, tc.LimitTransactionsToSave, tc.ExpectedError)
		})
	}
}

func testValidateTransactions(t *testing.T, tcName, trsFile, expTrsSaveResFile string, limTrs int, expErr error) {
	domain.LimitTransactionsToSave = limTrs
	var trs domain.Transactions
	domain.ReadJSON(t, trsFile, &trs)
	gotTrsSaveRes, gotErr := trs.ValidateTransactionsToSave()
	var expTrsSaveRes []domain.TransactionSaveResult

	if *update {
		domain.CreateJSON(t, expTrsSaveResFile, gotTrsSaveRes)
		return
	}

	domain.ReadJSON(t, expTrsSaveResFile, &expTrsSaveRes)

	domain.Compare(t, "compare transactions result", expTrsSaveRes, gotTrsSaveRes,
		cmpopts.SortSlices(func(i domain.TransactionSaveResult, j domain.TransactionSaveResult) bool {
			return i.Transaction.Description < j.Transaction.Description
		}))
	assert.Equal(t, expErr, gotErr, "expected error should be equal of got error")
}
