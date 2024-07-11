-- name: CreateTicket :one
INSERT INTO tickets (
  ticket_type_id
) VALUES (
  $1
)
RETURNING *;


-- name: GetTicketsByOrderID :many
SELECT * FROM ticket_order_items it
JOIN tickets t
ON t.ticket_id = it.ticket_id
JOIN ticket_types tty
ON tty.ticket_type_id=t.ticket_type_id
JOIN events ev
ON ev.event_id =tty.event_id
WHERE it.order_id = $1;
