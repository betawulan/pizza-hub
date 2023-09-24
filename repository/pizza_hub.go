package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/pizza-hub/model"
)

type pizzaHubRepo struct {
	db *sql.DB
}

func (p pizzaHubRepo) AddChef(ctx context.Context, chef model.Chef) (model.Chef, error) {

	query, args, err := sq.Insert("chef").
		Columns("name").
		Values(chef.Name).
		ToSql()
	if err != nil {
		return model.Chef{}, err
	}

	res, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Chef{}, err
	}

	chef.ID, err = res.LastInsertId()
	if err != nil {
		return model.Chef{}, err
	}

	return chef, nil
}

func (p pizzaHubRepo) GetMenus(ctx context.Context) ([]model.Menu, error) {

	query, args, err := sq.Select("id",
		"name").
		From("menu").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
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

func NewPizzaHubRepository(db *sql.DB) PizzaHubRepository {
	return pizzaHubRepo{
		db: db,
	}
}
