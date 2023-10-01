package main

import (
	"github.com/dragtor/tavern/domain/product"
	"github.com/dragtor/tavern/services/order"
	"github.com/dragtor/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		// order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}
	uid, err := os.AddCustomer("shubham")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	mac, err := product.NewProduct("macbook pro", "Laptop", 500)
	if err != nil {
		panic(err)
	}
	pod, err := product.NewProduct("airpod", "headset", 100)
	if err != nil {
		panic(err)
	}
	watch, err := product.NewProduct("iwatch", "watch", 70)
	if err != nil {
		panic(err)
	}
	return []product.Product{
		mac, pod, watch,
	}
}
