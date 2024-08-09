
-- name: CreateAttendee :one
INSERT INTO attendees (
  user_id
) VALUES (
  $1
)
RETURNING *;

-- name: GetAttendeeByUserID :one
SELECT * FROM attendees a
JOIN users u
ON a.user_id = u.user_id
WHERE a.attendee_id=$1
LIMIT 1; 

-- name: GetAttendee :one
SELECT * FROM attendees
WHERE attendee_id =$1
LIMIT 1;