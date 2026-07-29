-- name: CreateUser :one
INSERT INTO users_sqlc_demo (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password;

-- name: ListUsers :many
SELECT id, name, email, password
FROM users_sqlc_demo
ORDER BY id;

-- name: GetUser :one
SELECT id, name, email, password
FROM users_sqlc_demo
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users_sqlc_demo
SET name = $2,
	email = $3,
	password = $4
WHERE id = $1
RETURNING id, name, email, password;

-- name: DeleteUser :execrows
DELETE FROM users_sqlc_demo
WHERE id = $1;
