-- name: GetLocation :one
SELECT * FROM location
WHERE location_id = $1 LIMIT 1;

-- name: ListLocations :many
SELECT * FROM location
ORDER BY name;

-- name: CreateLocation :one
INSERT INTO location (
  name
) VALUES (
  $1
)
RETURNING *;

