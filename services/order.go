package services

import (
	"log"

	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/dragtor/impl-ddd/domain/customer"
	"github.com/dragtor/impl-ddd/domain/customer/memory"
	"github.com/dragtor/impl-ddd/domain/product"
	prodmem "github.com/dragtor/impl-ddd/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cgfs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cgfs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, product := range products {
			if err := pr.Add(product); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	cust, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	var products []aggregate.Product
	var total float64
	for _, id := range productIDs {
		prod, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, prod)
		total += prod.GetPrice()
	}
	log.Printf("Customer : %s has ordered %d products", cust.GetName(), len(products))
	return total, nil
}
