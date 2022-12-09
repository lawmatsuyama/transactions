package domain

import "github.com/google/uuid"

type UUIDRandomV4 struct{}

func NewUUIDRandomV4() UUIDRandomV4 {
	return UUIDRandomV4{}
}

func (random UUIDRandomV4) Generate() string {
	return uuid.New().String()
}

var (
	UUID UUIDGenerator = NewUUIDRandomV4()
)
