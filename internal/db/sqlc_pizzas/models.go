// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package pizzas_db

import ()

type Pizza struct {
	PizzaID     int32   `json:"pizzaId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
