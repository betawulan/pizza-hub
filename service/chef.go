package service

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository"
)

type chefService struct {
	chefRepo repository.ChefRepository
}

func (c chefService) AddChef(ctx context.Context, chef model.Chef) (model.Chef, error) {

	res, err := c.chefRepo.AddChef(ctx, chef)
	if err != nil {
		return model.Chef{}, err
	}

	return res, nil
}

func NewChefService(chefRepo repository.ChefRepository) ChefService {
	return chefService{
		chefRepo: chefRepo,
	}
}
