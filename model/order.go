package model

type PizzaType string

const (
	PizzaCheese PizzaType = "cheese"
	PizzaBBQ    PizzaType = "bbq"
)

type Order struct {
	ID        int64     `json:"id"`
	Customer  string    `json:"customer"`
	PizzaType PizzaType `json:"pizza_type"`
}
