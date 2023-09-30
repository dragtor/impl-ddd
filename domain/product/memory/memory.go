package memory

import (
	"sync"

	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, aggregate.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(product aggregate.Product) error {
	if _, ok := mpr.products[product.GetID()]; ok {
		return aggregate.ErrProductAlreadyExist
	}
	mpr.Lock()
	mpr.products[product.GetID()] = product
	mpr.Unlock()
	return nil
}

func (mpr *MemoryProductRepository) Update(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[product.GetID()]; ok {
		return aggregate.ErrProductAlreadyExist
	}

	mpr.products[product.GetID()] = product
	return nil
}

func (mpr *MemoryProductRepository) Delete(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[product.GetID()]; ok {
		return aggregate.ErrProductAlreadyExist
	}
	delete(mpr.products, product.GetID())
	return nil
}
