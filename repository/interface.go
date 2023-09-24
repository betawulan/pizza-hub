package repository

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
)

type PizzaHubRepository interface {
	AddChef(ctx context.Context, chef model.Chef) (model.Chef, error)
}
