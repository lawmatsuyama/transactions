package domain

import "context"

// FuncDBSession represents a function to open session with database
type FuncDBSession func(ctx context.Context) error
