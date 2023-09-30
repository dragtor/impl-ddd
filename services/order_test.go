package services

import (
	"testing"

	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	mac, err := aggregate.NewProduct("macbook pro", "Laptop", 500)
	if err != nil {
		t.Fatal(err)
	}
	pod, err := aggregate.NewProduct("airpod", "headset", 100)
	if err != nil {
		t.Fatal(err)
	}
	watch, err := aggregate.NewProduct("iwatch", "watch", 70)
	if err != nil {
		t.Fatal(err)
	}
	return []aggregate.Product{
		mac, pod, watch,
	}
}

func TestOrder_NewOrder(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	cust, err := aggregate.NewCustomer("shubham")
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
