package apimanager_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/infra/apimanager"
	"github.com/stretchr/testify/assert"
)

// InputTestSupport is a support struct to catch the input of functions
type InputTestSupport struct {
	UserID       string              `json:"user_id,omitempty"`
	Transactions domain.Transactions `json:"transactions,omitempty"`
}

func TestSave(t *testing.T) {
	testCases := []struct {
		Name                                  string
		TransactionsRequestFile               string
		TransactionsResultFile                string
		ExpectedInputSaveFile                 string
		ExpectedTransactionValidateResultFile string
		ErrSave                               error
		ErrWrite                              error
		ExpStatusCode                         int
	}{
		{
			Name:                                  "01_should_save_successfully_and_return_nil_error",
			TransactionsRequestFile:               "./testdata/transaction_app/01_should_save_successfully_and_return_nil_error/transactions.json",
			TransactionsResultFile:                "./testdata/transaction_app/01_should_save_successfully_and_return_nil_error/transactions_result.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_app/01_should_save_successfully_and_return_nil_error/exp_in_save.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_app/01_should_save_successfully_and_return_nil_error/exp_transactions_result.json",
			ExpStatusCode:                         http.StatusOK,
		},
		{
			Name:                                  "02_should_save_get_error_and_return_error",
			TransactionsRequestFile:               "./testdata/transaction_app/02_should_save_get_error_and_return_error/transactions.json",
			TransactionsResultFile:                "./testdata/transaction_app/02_should_save_get_error_and_return_error/transactions_result.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_app/02_should_save_get_error_and_return_error/exp_in_save.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_app/02_should_save_get_error_and_return_error/exp_transactions_result.json",
			ExpStatusCode:                         http.StatusBadRequest,
			ErrSave:                               domain.ErrInvalidTransaction,
		},
		{
			Name:                                  "03_should_write_get_error_and_return_error",
			TransactionsRequestFile:               "./testdata/transaction_app/03_should_write_get_error_and_return_error/transactions.json",
			TransactionsResultFile:                "./testdata/transaction_app/03_should_write_get_error_and_return_error/transactions_result.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_app/03_should_write_get_error_and_return_error/exp_in_save.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_app/03_should_write_get_error_and_return_error/exp_transactions_result.json",
			ErrWrite:                              domain.ErrUnknow,
			ExpStatusCode:                         http.StatusBadRequest,
		},
		{
			Name:                                  "04_should_json_decode_get_error_and_return_error",
			TransactionsRequestFile:               "./testdata/transaction_app/04_should_json_decode_get_error_and_return_error/transactions.json",
			TransactionsResultFile:                "./testdata/transaction_app/04_should_json_decode_get_error_and_return_error/transactions_result.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_app/04_should_json_decode_get_error_and_return_error/exp_in_save.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_app/04_should_json_decode_get_error_and_return_error/exp_transactions_result.json",
			ExpStatusCode:                         http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testSave(t, tc.Name, tc.TransactionsRequestFile, tc.TransactionsResultFile, tc.ExpectedInputSaveFile, tc.ExpectedTransactionValidateResultFile, tc.ErrSave, tc.ErrWrite, tc.ExpStatusCode)
		})
	}
}

func testSave(t *testing.T, name, trsFile, trsResultFile, expInSaveFile, expTrsResultFile string, errSave, errWrite error, expStatusCode int) {
	apimanager.Now = time.Date(2022, 10, 1, 10, 0, 0, 0, time.UTC)
	var gotInputSave InputTestSupport
	SaveMock = func(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error) {
		gotInputSave = InputTestSupport{UserID: userID, Transactions: transactions}
		var trsResult []domain.TransactionValidateResult
		domain.ReadJSON(t, trsResultFile, &trsResult)
		return trsResult, errSave
	}

	var gotStatusCode int
	WriteHeaderMock = func(statusCode int) {
		gotStatusCode = statusCode
	}

	var gotTrsResult apimanager.GenericResponse
	WriteMock = func(b []byte) (int, error) {
		err := json.Unmarshal(b, &gotTrsResult)
		if err != nil {
			return 0, nil
		}

		return len(b), errWrite
	}

	bTrs := domain.Read(t, trsFile)
	req, err := http.NewRequest("POST", "localhost:8888", bytes.NewBuffer(bTrs))
	if err != nil {
		t.Fatal(err)
	}

	transactionAPI := apimanager.NewTransactionAPI(mock{})
	transactionAPI.Save(mock{}, req)

	if *update {
		domain.CreateJSON(t, expInSaveFile, gotInputSave)
		domain.CreateJSON(t, expTrsResultFile, gotTrsResult)
		return
	}

	var expInSave InputTestSupport
	domain.ReadJSON(t, expInSaveFile, &expInSave)

	var expTrsResult apimanager.GenericResponse
	domain.ReadJSON(t, expTrsResultFile, &expTrsResult)

	assert.Equal(t, expStatusCode, gotStatusCode)
	domain.Compare(t, "compare input Save", expInSave, gotInputSave)
	domain.Compare(t, "compare input slice TransactionSaveResponse", expTrsResult.Result, gotTrsResult.Result,
		cmpopts.SortSlices(func(i, j apimanager.TransactionSaveResponse) bool {
			return i.Transaction.Description < j.Transaction.Description
		}))
	assert.Equal(t, expTrsResult.Error, gotTrsResult.Error, "expected error result should be equal got error result")
}
