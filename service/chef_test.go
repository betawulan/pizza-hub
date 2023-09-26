package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/betawulan/pizza-hub/model"
	"github.com/betawulan/pizza-hub/repository/mocks"
)

func Test_ChefService_Chef(t *testing.T) {
	tests := []struct {
		name   string
		ctx    context.Context
		req    model.Chef
		resp   model.Chef
		expErr error
	}{
		{
			name: "success",
			ctx:  context.Background(),
			req: model.Chef{
				ID:   1,
				Name: "wulan",
			},
			resp: model.Chef{
				ID:   1,
				Name: "wulan",
			},
			expErr: nil,
		},
		{
			name: "error",
			ctx:  context.Background(),
			req: model.Chef{
				ID:   1,
				Name: "wulan",
			},
			resp:   model.Chef{},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chefRepoMock := new(mocks.ChefRepository)

			chefRepoMock.On("AddChef", test.ctx, test.req).
				Return(test.resp, test.expErr).
				Once()

			chefService := NewChefService(chefRepoMock)

			response, err := chefService.AddChef(test.ctx, test.req)
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
