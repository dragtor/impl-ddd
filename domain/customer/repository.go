package customer

import (
	"errors"

	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCutomerNotFound     = errors.New("customer not found in repository")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
