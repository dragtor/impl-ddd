package order

import (
	"testing"

	"github.com/dragtor/tavern/domain/customer"
	"github.com/dragtor/tavern/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	mac, err := product.NewProduct("macbook pro", "Laptop", 500)
	if err != nil {
		t.Fatal(err)
	}
	pod, err := product.NewProduct("airpod", "headset", 100)
	if err != nil {
		t.Fatal(err)
	}
	watch, err := product.NewProduct("iwatch", "watch", 70)
	if err != nil {
		t.Fatal(err)
	}
	return []product.Product{
		mac, pod, watch,
	}
}

func TestOrder_NewOrder(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		// WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	cust, err := customer.NewCustomer("shubham")
	if err != nil {
		t.Fatal(err)
	}
	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
