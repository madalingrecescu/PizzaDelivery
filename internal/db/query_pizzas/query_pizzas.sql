-- name: CreatePizza :one
INSERT INTO pizzas (
    name,
    description,
    price

) VALUES (
             $1, $2, $3
         ) RETURNING *;

-- name: GetAllPizzas :many
SELECT * FROM pizzas
Order By id;

-- name: GetPizzaById :one
SELECT * FROM pizzas
WHERE pizza_id = $1 LIMIT 1;

-- name: GetPizzaByName :one
SELECT * FROM pizzas
WHERE name = $1 LIMIT 1;

-- name: