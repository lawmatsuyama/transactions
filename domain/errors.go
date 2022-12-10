package domain

import (
	"errors"
	"net/http"
)

// ErrorTransaction represents an error and implements error interface
type ErrorTransaction struct {
	ErrorOrigin error
	StatusCode  int
}

// Error return a string error
func (err ErrorTransaction) Error() string {
	return err.ErrorOrigin.Error()
}

// Error return an http status code
func (err ErrorTransaction) Status() int {
	return err.StatusCode
}

// ErrorToErrorTransaction converts error into ErrorTransaction
func ErrorToErrorTransaction(err error) ErrorTransaction {
	errTr, ok := err.(ErrorTransaction)
	if ok {
		return errTr
	}

	return ErrUnknow
}

var (
	ErrInvalidUser           = ErrorTransaction{ErrorOrigin: errors.New("invalid user"), StatusCode: http.StatusBadRequest}
	ErrDisabledUser          = ErrorTransaction{ErrorOrigin: errors.New("disabled user"), StatusCode: http.StatusBadRequest}
	ErrTransactionZeroAmount = ErrorTransaction{ErrorOrigin: errors.New("transaction amount is zero"), StatusCode: http.StatusBadRequest}
	ErrInvalidOperationType  = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction operation type"), StatusCode: http.StatusBadRequest}
	ErrInvalidOriginChannel  = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction origin channel"), StatusCode: http.StatusBadRequest}
	ErrInvalidTransaction    = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction"), StatusCode: http.StatusBadRequest}
	ErrUserNotFound          = ErrorTransaction{ErrorOrigin: errors.New("user not found"), StatusCode: http.StatusNotFound}
	ErrTransactionsNotFound  = ErrorTransaction{ErrorOrigin: errors.New("transactions not found"), StatusCode: http.StatusNotFound}
	ErrUnknow                = ErrorTransaction{ErrorOrigin: errors.New("unknow error"), StatusCode: http.StatusBadRequest}
)
