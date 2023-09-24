package service

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository"
)

type pizzaHubService struct {
	pizzaHubRepo repository.PizzaHubRepository
}

func (p pizzaHubService) AddChef(ctx context.Context, chef model.Chef) (model.Chef, error) {

	res, err := p.pizzaHubRepo.AddChef(ctx, chef)
	if err != nil {
		return model.Chef{}, err
	}

	return res, nil
}

func (p pizzaHubService) GetMenus(ctx context.Context) ([]model.Menu, error) {

	menus, err := p.pizzaHubRepo.GetMenus(ctx)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func NewPizzaHubService(pizzaHubRepo repository.PizzaHubRepository) PizzaHubService {
	return pizzaHubService{
		pizzaHubRepo: pizzaHubRepo,
	}
}
