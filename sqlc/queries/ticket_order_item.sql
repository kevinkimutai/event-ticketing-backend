
-- name: CreateTicketOrderItem :one
INSERT INTO ticket_order_items (
  order_id,ticket_id,quantity,total_price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTicketOrderItemByTicketID :one
SELECT * FROM ticket_order_items
WHERE ticket_id = $1 
LIMIT 1;

