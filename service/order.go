package service

import (
	"context"
	"time"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func (o orderService) Order(ctx context.Context, order model.Order) (model.Order, error) {

	go processOrder(order)

	res, err := o.orderRepo.Order(ctx, order)
	if err != nil {
		return model.Order{}, err
	}

	return res, nil
}

func processOrder(order model.Order) {
	var duration time.Duration

	switch order.PizzaType {
	case model.PizzaCheese:
		duration = 3 * time.Second
	case model.PizzaBBQ:
		duration = 5 * time.Second
	}

	time.Sleep(duration)
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return orderService{
		orderRepo: orderRepo,
	}
}
