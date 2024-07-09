-- name: CreateTicket :one
INSERT INTO tickets (
  ticket_type_id
) VALUES (
  $1
)
RETURNING *;

