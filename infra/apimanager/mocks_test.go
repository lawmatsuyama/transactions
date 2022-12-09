package apimanager_test

import (
	"context"
	"net/http"

	"github.com/lawmatsuyama/transactions/domain"
)

var (
	SaveMock        func(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error)
	WriteMock       func([]byte) (int, error)
	WriteHeaderMock func(statusCode int)
)

type mock struct{}

func (m mock) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionValidateResult, error) {
	return SaveMock(ctx, userID, transactions)
}

func (m mock) Header() http.Header {
	return make(http.Header)
}

func (m mock) Write(b []byte) (int, error) {
	return WriteMock(b)
}
func (m mock) WriteHeader(statusCode int) {
	WriteHeaderMock(statusCode)
}
