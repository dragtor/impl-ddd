package customer_test

import (
	"errors"
	"testing"

	"github.com/dragtor/tavern/domain/customer"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Customer with empty name",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		{
			test:        "Customer with valid name",
			name:        "shubham",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v , got %v \n", tc.expectedErr, err)
			}
		})
	}
}
