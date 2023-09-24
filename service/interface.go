package service

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
)

type PizzaHubService interface {
	AddChef(ctx context.Context, chef model.Chef) (model.Chef, error)
	GetMenus(ctx context.Context) ([]model.Menu, error)
}
