package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/betawulan/pizza-hub/model"
)

func Test_ChefRepository_Chef(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query, _, err := sq.Insert("chef").
		Columns("name").
		Values("ani").
		ToSql()
	require.NoError(t, err)

	query = regexp.QuoteMeta(query)

	type field struct {
		db *sql.DB
	}

	type args struct {
		ctx  context.Context
		chef model.Chef
	}

	tests := []struct {
		name    string
		field   field
		args    args
		expResp model.Chef
		expErr  error
		mockFn  func()
	}{
		{
			name:  "success",
			field: field{db: db},
			args: args{
				ctx: context.Background(),
				chef: model.Chef{
					ID:   1,
					Name: "ani"},
			},
			expResp: model.Chef{
				ID:   1,
				Name: "ani"},
			expErr: nil,
			mockFn: func() {
				mock.ExpectExec(query).
					WithArgs("ani").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chefRepoMock := NewChefRepository(db)

			test.mockFn()

			response, err := chefRepoMock.AddChef(test.args.ctx, test.args.chef)
			if err != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expErr, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expResp, response)
		})
	}

}
