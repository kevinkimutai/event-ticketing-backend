
-- name: CreateTicketOrder :one
INSERT INTO ticket_orders (
  attendee_id
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateTotalAmountOrder :exec
UPDATE ticket_orders 
  set total_amount = $2
  
WHERE order_id = $1;

-- name: GetTicketOrders :many
SELECT * FROM ticket_orders
LIMIT $1 OFFSET $2;
