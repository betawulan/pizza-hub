package service

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
)

type ChefService interface {
	AddChef(ctx context.Context, chef model.Chef) (model.Chef, error)
}

type MenuService interface {
	GetMenus(ctx context.Context) ([]model.Menu, error)
}
