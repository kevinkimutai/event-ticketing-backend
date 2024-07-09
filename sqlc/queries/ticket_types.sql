
-- name: CreateTicketType :one
INSERT INTO ticket_types (
  name,price,total_tickets,remaining_tickets,event_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTicketType :one
SELECT * FROM ticket_types
WHERE ticket_type_id = $1 LIMIT 1;

-- name: GetEventTicketTypes :many
SELECT * FROM ticket_types
WHERE event_id = $1;


-- name: UpdateRemainingTicketType :exec
UPDATE ticket_types
  set remaining_tickets = $2

WHERE ticket_type_id = $1;