
-- name: CreateOrganiser :one
INSERT INTO organisers (
  user_id,event_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetOrganisersByUserID :many
SELECT * FROM organisers 
WHERE user_id = $1;

-- name: GetCountOrganisersByUserID :many
SELECT COUNT(*) FROM organisers
WHERE user_id = $1;