package order

import (
	"context"
	"log"

	"github.com/dragtor/tavern/domain/customer"
	"github.com/dragtor/tavern/domain/customer/memory"
	"github.com/dragtor/tavern/domain/customer/mongo"
	"github.com/dragtor/tavern/domain/product"
	prodmem "github.com/dragtor/tavern/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
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

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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
	var products []product.Product
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

func (os *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	err = os.customers.Add(c)
	if err != nil {
		return uuid.Nil, nil
	}
	return c.GetID(), nil
}
