package usecases_test

import (
	"context"
	"testing"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/lawmatsuyama/transactions/usecases"
	"gotest.tools/assert"
)

// InputTestSupport is a support struct to catch the input of functions
type InputTestSupport struct {
	UserID       string                `json:"user_id,omitempty"`
	ExchangeName string                `json:"exchange_name,omitempty"`
	RoutingKey   string                `json:"routing_key,omitempty"`
	Transactions []*domain.Transaction `json:"transactions,omitempty"`
	Object       any                   `json:"object,omitempty"`
	Priority     uint8                 `json:"priority,omitempty"`
}

func TestSave(t *testing.T) {
	testCases := []struct {
		Name                                   string
		UserID                                 string
		TransactionsFiles                      string
		UserGetByIDFile                        string
		ExpectedInputUserGetByIDFile           string
		ExpectedInputSaveFile                  string
		ExpectedInputPublishFile               string
		ExepectedTransactionValidateResultFile string
		ErrUserGetByID                         error
		ErrSave                                error
		ErrPublish                             error
		ExpErr                                 error
	}{
		{
			Name:                                   "01_should_save_and_publish_transactions_and_return_nil_error",
			UserID:                                 "f982c57d-7632-40f5-b03e-a582f9a63e16",
			TransactionsFiles:                      "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/transactions.json",
			UserGetByIDFile:                        "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/user_get_by_id.json",
			ExpectedInputUserGetByIDFile:           "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/exp_in_user_get_by_id.json",
			ExpectedInputSaveFile:                  "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/exp_in_save.json",
			ExpectedInputPublishFile:               "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/exp_in_publish.json",
			ExepectedTransactionValidateResultFile: "./testdata/transaction_usecase/01_should_save_and_publish_transactions_and_return_nil_error/exp_transactions_result.json",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testSave(t,
				tc.Name,
				tc.UserID,
				tc.TransactionsFiles,
				tc.UserGetByIDFile,
				tc.ExpectedInputUserGetByIDFile,
				tc.ExpectedInputSaveFile,
				tc.ExpectedInputPublishFile,
				tc.ExepectedTransactionValidateResultFile,
				tc.ErrUserGetByID,
				tc.ErrSave,
				tc.ErrPublish,
				tc.ExpErr,
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
	SaveMock = func(ctx context.Context, transactions []*domain.Transaction) error {
		gotInputSave = InputTestSupport{Transactions: transactions}
		return errSave
	}

	var gotInputPub InputTestSupport
	PublishMock = func(ctx context.Context, excName, routingKey string, obj any, priority uint8) error {
		gotInputPub = InputTestSupport{ExchangeName: excName, RoutingKey: routingKey, Object: obj, Priority: priority}
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

	var expTrsResult []domain.TransactionValidateResult
	domain.ReadJSON(t, expTrsResultFile, &expTrsResult)

	assert.Equal(t, expErr, err)
	domain.Compare(t, "compare input user GetByID", expInUserGetByID, gotInUserGetByID)
	domain.Compare(t, "compare input Save", expInSave, gotInputSave)
	domain.Compare(t, "compare input Publish", expInPub, gotInputPub)
	domain.Compare(t, "compare input TransactionValidateResult", expTrsResult, gotTrsResult)

}
