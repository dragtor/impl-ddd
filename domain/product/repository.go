package product

import (
	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
}
