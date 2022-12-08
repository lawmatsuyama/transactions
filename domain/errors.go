package domain

import "errors"

var (
	ErrInvalidUser           = errors.New("invalid user")
	ErrDisabledUser          = errors.New("disabled user")
	ErrTransactionZeroAmount = errors.New("transaction amount is zero")
	ErrInvalidOperationType  = errors.New("invalid transaction operation type")
	ErrInvalidOriginChannel  = errors.New("invalid transaction origin channel")
	ErrInvalidTransaction    = errors.New("invalid transaction")
)
