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

func NewPizzaHubRepository(db *sql.DB) PizzaHubRepository {
	return pizzaHubRepo{
		db: db,
	}
}
