-- name: CreateUser :one
INSERT INTO users (
    name,
    email
) VALUES (
    sqlc.arg(name),
    sqlc.arg(email)
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET email = sqlc.arg(email)
WHERE id = sqlc.arg(id)
RETURNING *;
