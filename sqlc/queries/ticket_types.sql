
-- name: CreateTicketType :one
INSERT INTO ticket_types (
  name,price,total_tickets
) VALUES (
  $1, $2, $3
)
RETURNING *;

