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
Order By pizza_id;

-- name: GetPizzaById :one
SELECT * FROM pizzas
WHERE pizza_id = $1 LIMIT 1;

-- name: GetPizzaByName :one
SELECT * FROM pizzas
WHERE name = $1 LIMIT 1;

-- name: UpdatePizza :one
UPDATE pizzas
SET name = $2, description = $3, price = $4
WHERE pizza_id = $1
RETURNING *;

-- name: DeletePizza :exec
DELETE FROM pizzas
WHERE pizza_id = $1;