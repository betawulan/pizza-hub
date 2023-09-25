package repository

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
)

type ChefRepository interface {
	AddChef(ctx context.Context, chef model.Chef) (model.Chef, error)
}

type MenuRepository interface {
	GetMenus(ctx context.Context) ([]model.Menu, error)
}

type OrderRepository interface {
	Order(ctx context.Context, order model.Order) (model.Order, error)
}
