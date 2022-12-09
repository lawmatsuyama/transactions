package usecases_test

import "testing"

func TestSave(t *testing.T) {
	testCases := []struct {
		Name string
	}{
		{
			Name: "01_should_save_and_publish_transactions_and_return_nil_error",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

		})
	}
}

func testSave(t *testing.T) {

}
