package domain

import "github.com/google/uuid"

// UUIDRandomV4 generates a new UUID version 4
type UUIDRandomV4 struct{}

// NewUUIDRandomV4 returns UUIDRandomV4
func NewUUIDRandomV4() UUIDRandomV4 {
	return UUIDRandomV4{}
}

// Generate returns a new string UUID
func (random UUIDRandomV4) Generate() string {
	return uuid.New().String()
}

var (
	// UUID is a global variable to use as UUID generator.
	//It usefull to set another uuid generator in unit tests
	UUID UUIDGenerator = NewUUIDRandomV4()
)
