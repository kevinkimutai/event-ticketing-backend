-- name: CreateUser :one
INSERT INTO users (
  full_name,email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUserID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;