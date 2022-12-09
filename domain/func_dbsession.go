package domain

import "context"

type FuncDBSession func(ctx context.Context) error
