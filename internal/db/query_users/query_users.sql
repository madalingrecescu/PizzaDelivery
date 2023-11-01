-- name: CreateAccount :one
INSERT INTO users (
    username,
    email,
    hashed_password,
    phone_number

) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
