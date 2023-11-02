// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query_pizzas.sql

package pizzas_db

import (
	"context"
)

const createPizza = `-- name: CreatePizza :one
INSERT INTO pizzas (
    name,
    description,
    price

) VALUES (
             $1, $2, $3
         ) RETURNING pizza_id, name, description, price
`

type CreatePizzaParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (q *Queries) CreatePizza(ctx context.Context, arg CreatePizzaParams) (Pizza, error) {
	row := q.db.QueryRowContext(ctx, createPizza, arg.Name, arg.Description, arg.Price)
	var i Pizza
	err := row.Scan(
		&i.PizzaID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const deletePizza = `-- name: DeletePizza :exec
DELETE FROM pizzas
WHERE pizza_id = $1
`

func (q *Queries) DeletePizza(ctx context.Context, pizzaID int32) error {
	_, err := q.db.ExecContext(ctx, deletePizza, pizzaID)
	return err
}

const getAllPizzas = `-- name: GetAllPizzas :many
SELECT pizza_id, name, description, price FROM pizzas
Order By pizza_id
`

func (q *Queries) GetAllPizzas(ctx context.Context) ([]Pizza, error) {
	rows, err := q.db.QueryContext(ctx, getAllPizzas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pizza
	for rows.Next() {
		var i Pizza
		if err := rows.Scan(
			&i.PizzaID,
			&i.Name,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPizzaById = `-- name: GetPizzaById :one
SELECT pizza_id, name, description, price FROM pizzas
WHERE pizza_id = $1 LIMIT 1
`

func (q *Queries) GetPizzaById(ctx context.Context, pizzaID int32) (Pizza, error) {
	row := q.db.QueryRowContext(ctx, getPizzaById, pizzaID)
	var i Pizza
	err := row.Scan(
		&i.PizzaID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getPizzaByName = `-- name: GetPizzaByName :one
SELECT pizza_id, name, description, price FROM pizzas
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPizzaByName(ctx context.Context, name string) (Pizza, error) {
	row := q.db.QueryRowContext(ctx, getPizzaByName, name)
	var i Pizza
	err := row.Scan(
		&i.PizzaID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const updatePizza = `-- name: UpdatePizza :one
UPDATE pizzas
SET name = $2, description = $3, price = $4
WHERE pizza_id = $1
RETURNING pizza_id, name, description, price
`

type UpdatePizzaParams struct {
	PizzaID     int32   `json:"pizzaId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (q *Queries) UpdatePizza(ctx context.Context, arg UpdatePizzaParams) (Pizza, error) {
	row := q.db.QueryRowContext(ctx, updatePizza,
		arg.PizzaID,
		arg.Name,
		arg.Description,
		arg.Price,
	)
	var i Pizza
	err := row.Scan(
		&i.PizzaID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}
