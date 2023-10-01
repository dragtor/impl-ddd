package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCutomerNotFound     = errors.New("customer not found in repository")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed to update customer")
)

type Repository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
