package domain_test

import (
	"testing"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserIsValid(t *testing.T) {
	testCases := []struct {
		Name          string
		UserFile      string
		ExpectedError error
	}{
		{
			Name:          "01_should_user_is_valid_return_nil_error",
			UserFile:      "./testdata/user/01_should_user_is_valid_return_nil_error/user.json",
			ExpectedError: nil,
		},
		{
			Name:          "02_should_user_is_valid_return_disabled_user_error",
			UserFile:      "./testdata/user/02_should_user_is_valid_return_disabled_user_error/user.json",
			ExpectedError: domain.ErrDisabledUser,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testUserIsValid(t, tc.Name, tc.UserFile, tc.ExpectedError)
		})
	}
}

func testUserIsValid(t *testing.T, tcName, userFile string, exp error) {
	var user domain.User
	domain.ReadJSON(t, userFile, &user)
	got := user.IsValid()
	assert.Equal(t, exp, got, "expected error should be equal of got error")
}
