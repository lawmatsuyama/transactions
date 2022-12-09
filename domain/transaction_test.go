package domain_test

import (
	"testing"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	testCases := []struct {
		Name            string
		TransactionFile string
		ExpectedErrors  []string
	}{
		{
			Name:            "01_should_return_empty_errors",
			TransactionFile: "./testdata/transaction/01_should_return_empty_errors/transaction.json",
			ExpectedErrors:  []string{},
		},
		{
			Name:            "02_should_return_three_errors",
			TransactionFile: "./testdata/transaction/02_should_return_three_errors/transaction.json",
			ExpectedErrors: []string{
				"transaction amount is zero",
				"invalid transaction operation type",
				"invalid transaction origin channel",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testTransactionIsValid(t, tc.Name, tc.TransactionFile, tc.ExpectedErrors)
		})
	}
}

func testTransactionIsValid(t *testing.T, tcName, trFile string, expErrs []string) {
	var tr domain.Transaction
	domain.ReadJSON(t, trFile, &tr)

	gotErrs := tr.IsValid()
	assert.Equal(t, expErrs, gotErrs, "expected errors should be equal of got errors")
}
