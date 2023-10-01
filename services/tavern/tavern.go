package tavern

import (
	"log"

	"github.com/dragtor/tavern/services/order"
	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService *order.OrderService
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	ts := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(ts)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("\nBill custemer of price %0.0f\n", price)
	return nil
}
