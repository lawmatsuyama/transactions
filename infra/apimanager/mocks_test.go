package apimanager_test

import (
	"context"
	"net/http"

	"github.com/lawmatsuyama/transactions/domain"
)

var (
	SaveMock        func(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionSaveResult, error)
	GetMock         func(ctx context.Context, filter domain.TransactionFilter) (domain.TransactionsPaging, error)
	WriteMock       func([]byte) (int, error)
	WriteHeaderMock func(statusCode int)
)

type mock struct{}

func (m mock) Save(ctx context.Context, userID string, transactions domain.Transactions) ([]domain.TransactionSaveResult, error) {
	return SaveMock(ctx, userID, transactions)
}
func (m mock) Get(ctx context.Context, filter domain.TransactionFilter) (domain.TransactionsPaging, error) {
	return GetMock(ctx, filter)
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
