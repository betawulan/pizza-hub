package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/pizza-hub/model"
)

type menuRepo struct {
	db *sql.DB
}

func (m menuRepo) GetMenus(ctx context.Context) ([]model.Menu, error) {

	query, args, err := sq.Select("id",
		"name").
		From("menu").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	menus := make([]model.Menu, 0)
	for rows.Next() {
		var menu model.Menu

		err = rows.Scan(&menu.ID,
			&menu.Name)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func NewMenuRepository(db *sql.DB) MenuRepository {
	return menuRepo{
		db: db,
	}
}
