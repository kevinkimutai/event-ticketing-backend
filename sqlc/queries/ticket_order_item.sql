
-- name: CreateTicketOrderItem :one
INSERT INTO ticket_order_items (
  order_id,ticket_id,quantity,total_price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;



