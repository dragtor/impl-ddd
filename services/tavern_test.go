package services

import (
	"testing"

	"github.com/dragtor/impl-ddd/aggregate"
	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	cust, err := aggregate.NewCustomer("shubham")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}
