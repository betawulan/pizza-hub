package service

import (
	"context"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository"
)

type menuService struct {
	menuRepo repository.MenuRepository
}

func (m menuService) GetMenus(ctx context.Context) ([]model.Menu, error) {

	menus, err := m.menuRepo.GetMenus(ctx)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return menuService{
		menuRepo: menuRepo,
	}
}
