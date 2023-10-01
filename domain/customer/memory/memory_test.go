package memory

import (
	"errors"
	"testing"

	"github.com/dragtor/tavern/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	cust, err := customer.NewCustomer("shubham")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}
	testcases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("7e9ac9c2-5a89-48ed-ab0c-0a19b10a0dd9"),
			expectedErr: customer.ErrCutomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected Error %v , Got %v", tc.expectedErr, err)
			}
		})
	}

}
