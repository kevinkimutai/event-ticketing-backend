
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
