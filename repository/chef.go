package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/pizza-hub/model"
)

type chefRepo struct {
	db *sql.DB
}

func (c chefRepo) AddChef(ctx context.Context, chef model.Chef) (model.Chef, error) {

	query, args, err := sq.Insert("chef").
		Columns("name").
		Values(chef.Name).
		ToSql()
	if err != nil {
		return model.Chef{}, err
	}

	res, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Chef{}, err
	}

	chef.ID, err = res.LastInsertId()
	if err != nil {
		return model.Chef{}, err
	}

	return chef, nil
}

func NewChefRepository(db *sql.DB) ChefRepository {
	return chefRepo{
		db: db,
	}
}
