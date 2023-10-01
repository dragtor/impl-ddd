package tavern

import (
	"testing"

	"github.com/dragtor/tavern/domain/product"
	"github.com/dragtor/tavern/services/order"
	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	uid, err := os.AddCustomer("shubham")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
		products[2].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}

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
