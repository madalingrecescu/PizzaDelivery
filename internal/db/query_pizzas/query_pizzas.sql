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

-- name: CreateShoppingCart :one
INSERT INTO shopping_cart (
    username
) VALUES (
          $1
         ) RETURNING *;

-- name: GetShoppingCartByUsername :one
SELECT *
FROM shopping_cart
WHERE username = $1
ORDER BY shopping_cart_id DESC
    LIMIT 1;


-- name: CreatePizzaOrder :one
INSERT INTO pizza_order (
    shopping_cart_id,
    pizza_name,
    pizza_price,
    quantity

) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: AddQuantityToOrderForExistingPizza :one
UPDATE pizza_order
SET quantity = quantity + 1
WHERE pizza_name = $1 AND shopping_cart_id = $2
    RETURNING *;

-- name: SubtractQuantityOfExistingPizzaFromOrder :one
UPDATE pizza_order
SET quantity = quantity - 1
WHERE pizza_name = $1 AND shopping_cart_id = $2
    RETURNING *;

-- name: GetAllOrdersByShoppingCartID :many
SELECT *
FROM pizza_order
WHERE shopping_cart_id = $1;


-- name: DeleteAllPizzasWithTheSameNameFromShoppingCart :exec
DELETE FROM pizza_order
WHERE shopping_cart_id = $1 AND pizza_name = $2;

-- name: GetPizzaOrderByNameFromShoppingCart :one
SELECT *
FROM pizza_order
WHERE shopping_cart_id = $1 AND pizza_name = $2 LIMIT 1;

-- name: UpdatePizzaQuantityInShoppingCart :exec
UPDATE pizza_order
SET quantity = $1
WHERE pizza_name = $2 AND shopping_cart_id = $3;