
-- name: CreateOrganiser :one
INSERT INTO organisers (
  user_id,event_id
) VALUES (
  $1, $2
)
RETURNING *;

