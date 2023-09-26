package service

import (
	"context"
	"errors"
	"testing"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository/mocks"
	"github.com/stretchr/testify/require"
)

func Test_MenuService_Menu(t *testing.T) {
	menu1 := model.Menu{
		ID:   1,
		Name: "pizza bbq",
	}

	menu2 := model.Menu{
		ID:   2,
		Name: "pizza cheese",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		resp   []model.Menu
		expErr error
	}{
		{
			name:   "success",
			ctx:    context.Background(),
			resp:   []model.Menu{menu1, menu2},
			expErr: nil,
		},
		{
			name:   "error",
			ctx:    context.Background(),
			resp:   []model.Menu{},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			menuMockRepo := new(mocks.MenuRepository)

			menuMockRepo.On("GetMenus", test.ctx).
				Return(test.resp, test.expErr).
				Once()

			menuService := NewMenuService(menuMockRepo)

			response, err := menuService.GetMenus(test.ctx)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.resp, response)
		})
	}

}
