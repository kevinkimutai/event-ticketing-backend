
-- name: CreateAttendee :one
INSERT INTO attendees (
  user_id
) VALUES (
  $1
)
RETURNING *;

