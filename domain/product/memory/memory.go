package memory

import (
	"sync"

	"github.com/dragtor/tavern/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(prod product.Product) error {
	if _, ok := mpr.products[prod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	mpr.Lock()
	mpr.products[prod.GetID()] = prod
	mpr.Unlock()
	return nil
}

func (mpr *MemoryProductRepository) Update(prod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[prod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[prod.GetID()] = prod
	return nil
}

func (mpr *MemoryProductRepository) Delete(prod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[prod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	delete(mpr.products, prod.GetID())
	return nil
}
