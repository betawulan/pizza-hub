package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/pizza-hub/model"
)

type orderRepo struct {
	db *sql.DB
}

func (o orderRepo) Order(ctx context.Context, order model.Order) (model.Order, error) {

	query, args, err := sq.Insert("orders").
		Columns("customer",
			"pizza_type").
		Values(order.Customer,
			order.PizzaType).
		ToSql()
	if err != nil {
		return model.Order{}, err
	}

	res, err := o.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.Order{}, err
	}

	order.ID, err = res.LastInsertId()
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return orderRepo{
		db: db,
	}
}
