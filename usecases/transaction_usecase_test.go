package usecases_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/usecases"
	"gotest.tools/assert"
)

// InputTestSupport is a support struct to catch the input of functions
type InputTestSupport struct {
	UserID       string                `json:"user_id,omitempty"`
	ExchangeName string                `json:"exchange_name,omitempty"`
	RoutingKey   string                `json:"routing_key,omitempty"`
	Object       domain.Transactions   `json:"object,omitempty"`
	Transactions []*domain.Transaction `json:"transactions,omitempty"`
	Priority     uint8                 `json:"priority,omitempty"`
}

func TestSave(t *testing.T) {
	testCases := []struct {
		Name                                  string
		UserID                                string
		TransactionsFile                      string
		UserGetByIDFile                       string
		ExpectedInputUserGetByIDFile          string
		ExpectedInputSaveFile                 string
		ExpectedInputPublishFile              string
		ExpectedTransactionValidateResultFile string
		ErrUserGetByID                        error
		ErrSave                               error
		ErrPublish                            error
		ExpectedError                         error
	}{
		{
			Name:                                  "01_should_save_and_publish_transactions_successfully_and_return_nil_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/01_should_save_and_publish_transactions_successfully_and_return_nil_error/exp_transactions_result.json",
		},
		{
			Name:                                  "02_should_not_found_user_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/02_should_not_found_user_and_return_error/exp_transactions_result.json",
			ErrUserGetByID:                        domain.ErrUserNotFound,
			ExpectedError:                         domain.ErrUserNotFound,
		},
		{
			Name:                                  "03_should_get_error_on_is_valid_user_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/03_should_trigger_invalid_user_and_return_error/exp_transactions_result.json",
			ExpectedError:                         domain.ErrDisabledUser,
		},
		{
			Name:                                  "04_should_get_error_on_validate_transactions_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/04_should_get_error_on_validate_transactions_and_return_error/exp_transactions_result.json",
			ExpectedError:                         domain.ErrInvalidTransaction,
		},
		{
			Name:                                  "05_should_get_error_on_save_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/05_should_get_error_on_save_and_return_error/exp_transactions_result.json",
			ErrSave:                               domain.ErrUnknow,
			ExpectedError:                         domain.ErrUnknow,
		},
		{
			Name:                                  "06_should_get_error_on_publish_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/06_should_get_error_on_publish_and_return_error/exp_transactions_result.json",
			ErrPublish:                            domain.ErrUnknow,
			ExpectedError:                         domain.ErrUnknow,
		},
		{
			Name:                                  "07_should_get_user_is_empty_and_return_error",
			UserID:                                "",
			TransactionsFile:                      "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/07_should_get_user_is_empty_and_return_error/exp_transactions_result.json",
			ExpectedError:                         domain.ErrInvalidUser,
		},
		{
			Name:                                  "08_should_get_no_transcations_request_and_return_error",
			UserID:                                "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFile:                      "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/transactions.json",
			UserGetByIDFile:                       "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:          "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                 "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/exp_in_save.json",
			ExpectedInputPublishFile:              "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/exp_in_publish.json",
			ExpectedTransactionValidateResultFile: "./testdata/transaction_usecase/save/08_should_get_no_transcations_request_and_return_error/exp_transactions_result.json",
			ExpectedError:                         domain.ErrInvalidTransaction,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testSave(t,
				tc.Name,
				tc.UserID,
				tc.TransactionsFile,
				tc.UserGetByIDFile,
				tc.ExpectedInputUserGetByIDFile,
				tc.ExpectedInputSaveFile,
				tc.ExpectedInputPublishFile,
				tc.ExpectedTransactionValidateResultFile,
				tc.ErrUserGetByID,
				tc.ErrSave,
				tc.ErrPublish,
				tc.ExpectedError,
			)
		})
	}
}

func testSave(t *testing.T, name, userID, trsFile, userGetByIDFile, expInUserGetByIDFile, expInSaveFile, expInPubFile, expTrsResultFile string, errGetByID, errSave, errPublish, expErr error) {
	var gotInUserGetByID InputTestSupport
	GetByIDMock = func(ctx context.Context, id string) (domain.User, error) {
		gotInUserGetByID = InputTestSupport{UserID: id}
		var user domain.User
		domain.ReadJSON(t, userGetByIDFile, &user)
		return user, errGetByID
	}

	var gotInputSave InputTestSupport
	SaveMock = func(ctx context.Context, transactions domain.Transactions) error {
		gotInputSave = InputTestSupport{Transactions: transactions}
		return errSave
	}

	var gotInputPub InputTestSupport
	PublishMock = func(ctx context.Context, excName, routingKey string, obj any, priority uint8) error {
		trs := obj.(domain.Transactions)
		gotInputPub = InputTestSupport{ExchangeName: excName, RoutingKey: routingKey, Object: trs, Priority: priority}
		return errPublish
	}

	WithSessionMock = func(ctx context.Context, f domain.FuncDBSession) error {
		return f(ctx)
	}

	var trs domain.Transactions
	domain.ReadJSON(t, trsFile, &trs)
	transactionUseCase := usecases.NewTransactionUseCase(mock{}, mock{}, mock{}, mock{})
	gotTrsResult, err := transactionUseCase.Save(context.Background(), userID, trs)

	if *update {
		domain.CreateJSON(t, expInUserGetByIDFile, gotInUserGetByID)
		domain.CreateJSON(t, expInSaveFile, gotInputSave)
		domain.CreateJSON(t, expInPubFile, gotInputPub)
		domain.CreateJSON(t, expTrsResultFile, gotTrsResult)
		return
	}

	var expInUserGetByID InputTestSupport
	domain.ReadJSON(t, expInUserGetByIDFile, &expInUserGetByID)

	var expInSave InputTestSupport
	domain.ReadJSON(t, expInSaveFile, &expInSave)

	var expInPub InputTestSupport
	domain.ReadJSON(t, expInPubFile, &expInPub)

	var expTrsResult []domain.TransactionSaveResult
	domain.ReadJSON(t, expTrsResultFile, &expTrsResult)

	assert.Equal(t, expErr, err)
	domain.Compare(t, "compare input user GetByID", expInUserGetByID, gotInUserGetByID)
	domain.Compare(t, "compare input Save", expInSave, gotInputSave)
	domain.Compare(t, "compare input Publish", expInPub, gotInputPub)
	domain.Compare(t, "compare input TransactionValidateResult", expTrsResult, gotTrsResult,
		cmpopts.SortSlices(func(i, j domain.TransactionSaveResult) bool {
			return i.Transaction.Description < j.Transaction.Description
		}))

}

func TestGet(t *testing.T) {
	testCases := []struct {
		Name                     string
		FilterFile               string
		TransactionsFile         string
		ExpectedInputGetFile     string
		ExpectedTransactionsFile string
		LimitTransactionsByPage  int64
		ErrorGet                 error
		ExpectedError            error
	}{
		{
			Name:                     "01_should_return_transactions_without_nextpage",
			FilterFile:               "./testdata/transaction_usecase/get/01_should_return_transactions_without_nextpage/filter.json",
			TransactionsFile:         "./testdata/transaction_usecase/get/01_should_return_transactions_without_nextpage/transactions.json",
			ExpectedInputGetFile:     "./testdata/transaction_usecase/get/01_should_return_transactions_without_nextpage/exp_in_get.json",
			ExpectedTransactionsFile: "./testdata/transaction_usecase/get/01_should_return_transactions_without_nextpage/exp_transactions.json",
			LimitTransactionsByPage:  20,
		},
		{
			Name:                     "02_should_return_filter_validate_error",
			FilterFile:               "./testdata/transaction_usecase/get/02_should_return_filter_validate_error/filter.json",
			TransactionsFile:         "./testdata/transaction_usecase/get/02_should_return_filter_validate_error/transactions.json",
			ExpectedInputGetFile:     "./testdata/transaction_usecase/get/02_should_return_filter_validate_error/exp_in_get.json",
			ExpectedTransactionsFile: "./testdata/transaction_usecase/get/02_should_return_filter_validate_error/exp_transactions.json",
			ExpectedError:            domain.ErrInvalidOriginChannel,
		},
		{
			Name:                     "03_get_repository_should_return_error",
			FilterFile:               "./testdata/transaction_usecase/get/03_get_repository_should_return_error/filter.json",
			TransactionsFile:         "./testdata/transaction_usecase/get/03_get_repository_should_return_error/transactions.json",
			ExpectedInputGetFile:     "./testdata/transaction_usecase/get/03_get_repository_should_return_error/exp_in_get.json",
			ExpectedTransactionsFile: "./testdata/transaction_usecase/get/03_get_repository_should_return_error/exp_transactions.json",
			ErrorGet:                 domain.ErrUnknow,
			ExpectedError:            domain.ErrUnknow,
		},
		{
			Name:                     "04_should_return_transactions_with_nextpage",
			FilterFile:               "./testdata/transaction_usecase/get/04_should_return_transactions_with_nextpage/filter.json",
			TransactionsFile:         "./testdata/transaction_usecase/get/04_should_return_transactions_with_nextpage/transactions.json",
			ExpectedInputGetFile:     "./testdata/transaction_usecase/get/04_should_return_transactions_with_nextpage/exp_in_get.json",
			ExpectedTransactionsFile: "./testdata/transaction_usecase/get/04_should_return_transactions_with_nextpage/exp_transactions.json",
			LimitTransactionsByPage:  2,
		},
		{
			Name:                     "05_no_transactions_should_return_transactions_error_not_found_transactions",
			FilterFile:               "./testdata/transaction_usecase/get/05_no_transactions_should_return_transactions_error_not_found_transactions/filter.json",
			TransactionsFile:         "./testdata/transaction_usecase/get/05_no_transactions_should_return_transactions_error_not_found_transactions/transactions.json",
			ExpectedInputGetFile:     "./testdata/transaction_usecase/get/05_no_transactions_should_return_transactions_error_not_found_transactions/exp_in_get.json",
			ExpectedTransactionsFile: "./testdata/transaction_usecase/get/05_no_transactions_should_return_transactions_error_not_found_transactions/exp_transactions.json",
			ExpectedError:            domain.ErrTransactionsNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testGet(t, tc.Name, tc.FilterFile, tc.TransactionsFile, tc.ExpectedInputGetFile, tc.ExpectedTransactionsFile, tc.LimitTransactionsByPage, tc.ErrorGet, tc.ExpectedError)
		})
	}
}

func testGet(t *testing.T, name string, filterFile, trsFile, expInGetFile, expTrsFile string, limitTrs int64, errGet, expErr error) {
	domain.LimitTransactionsByPage = limitTrs
	var filter domain.TransactionFilter
	domain.ReadJSON(t, filterFile, &filter)

	var gotInGet domain.TransactionFilter
	GetMock = func(ctx context.Context, filterTrs domain.TransactionFilter) (domain.Transactions, error) {
		gotInGet = filterTrs
		var trs domain.Transactions
		domain.ReadJSON(t, trsFile, &trs)
		return trs, errGet
	}

	transactionUseCase := usecases.NewTransactionUseCase(mock{}, mock{}, mock{}, mock{})
	gotTrs, err := transactionUseCase.Get(context.Background(), filter)
	if *update {
		domain.CreateJSON(t, expInGetFile, gotInGet)
		domain.CreateJSON(t, expTrsFile, gotTrs)
		return
	}

	var expInGet domain.TransactionFilter
	domain.ReadJSON(t, expInGetFile, &expInGet)

	var expTrs domain.TransactionsPaging
	domain.ReadJSON(t, expTrsFile, &expTrs)

	assert.Equal(t, expErr, err)
	domain.Compare(t, "compare input of Get func", expInGet, gotInGet)
	domain.Compare(t, "compare transactions", expTrs, gotTrs)
}
