package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"

	"github.com/betawulan/pizza-hub/model"
)

func Test_MenuReposiroty_Menu(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query, _, err := sq.Select("id",
		"name").
		From("menu").
		ToSql()
	require.NoError(t, err)

	query = regexp.QuoteMeta(query)

	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
	}

	menu1 := model.Menu{
		ID:   1,
		Name: "pizza cheese",
	}

	menu2 := model.Menu{
		ID:   2,
		Name: "pizza bbq",
	}

	respMenus := []model.Menu{menu1, menu2}
	_rowsMock := sqlmock.NewRows([]string{"id", "menu"}).AddRow(1, "pizza cheese").AddRow(2, "pizza bbq")

	tests := []struct {
		name    string
		fields  fields
		args    args
		expResp []model.Menu
		expErr  error
		mockFn  func()
	}{
		{
			name:    "success",
			fields:  fields{db: db},
			args:    args{ctx: context.Background()},
			expResp: respMenus,
			expErr:  nil,
			mockFn: func() {
				mockDB.ExpectQuery(query).WillReturnRows(_rowsMock)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			menuRepoMock := NewMenuRepository(db)

			test.mockFn()

			response, err := menuRepoMock.GetMenus(test.args.ctx)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expResp, response)
		})
	}
}