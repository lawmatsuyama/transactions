package domain_test

import (
	"testing"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateTransactions(t *testing.T) {
	testCases := []struct {
		Name                                   string
		TransactionsFile                       string
		ExepectedTransactionValidateResultFile string
		ExpectedError                          error
	}{
		{
			Name:                                   "01_should_validate_transactions_return_error_nil",
			TransactionsFile:                       "./testdata/_transactions/01_should_validate_transactions_return_error_nil/transactions.json",
			ExepectedTransactionValidateResultFile: "./testdata/_transactions/01_should_validate_transactions_return_error_nil/transactions_result.json",
			ExpectedError:                          nil,
		},
		{
			Name:                                   "02_should_validate_transactions_return_invalid_operation",
			TransactionsFile:                       "./testdata/_transactions/02_should_validate_transactions_return_invalid_operation/transactions.json",
			ExepectedTransactionValidateResultFile: "./testdata/_transactions/02_should_validate_transactions_return_invalid_operation/transactions_result.json",
			ExpectedError:                          domain.ErrInvalidTransaction,
		},
		{
			Name:                                   "03_should_validate_transactions_return_invalid_origin_and_operation",
			TransactionsFile:                       "./testdata/_transactions/03_should_validate_transactions_return_invalid_origin_and_operation/transactions.json",
			ExepectedTransactionValidateResultFile: "./testdata/_transactions/03_should_validate_transactions_return_invalid_origin_and_operation/transactions_result.json",
			ExpectedError:                          domain.ErrInvalidTransaction,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testValidateTransactions(t, tc.Name, tc.TransactionsFile, tc.ExepectedTransactionValidateResultFile, tc.ExpectedError)
		})
	}
}

func testValidateTransactions(t *testing.T, tcName, trsFile, expTrsValResFile string, expErr error) {
	var trs domain.Transactions
	domain.ReadJSON(t, trsFile, &trs)
	gotTrsValRes, gotErr := trs.ValidateTransactions()
	var expTrsValRes []domain.TransactionValidateResult

	if *update {
		domain.CreateJSON(t, expTrsValResFile, gotTrsValRes)
		return
	}

	domain.ReadJSON(t, expTrsValResFile, &expTrsValRes)

	domain.Compare(t, "compare transactions result", expTrsValRes, gotTrsValRes)
	assert.Equal(t, expErr, gotErr, "expected error should be equal of got error")
}
