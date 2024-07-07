-- name: GetCategory :one
SELECT * FROM categories
WHERE category_id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name;

-- name: CreateCategory :one
INSERT INTO categories (
  name
) VALUES (
  $1
)
RETURNING *;

