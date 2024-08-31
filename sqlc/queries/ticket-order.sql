
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

-- name: GetTicketOrder :one
SELECT * FROM ticket_orders
WHERE order_id =$1
LIMIT 1;

-- name: GetTicketOrderDetails :one
SELECT 
	tord.order_id AS order_id,
	u.full_name AS full_name,
	items.quantity AS quantity,
	ttypes.name AS ticket_type_name,
	e.name AS event_name,
	e.date AS event_date,
	e.location AS event_location,
	tord.admit_status AS admitted
FROM ticket_order_items items
JOIN tickets t ON t.ticket_id = items.ticket_id
JOIN ticket_types ttypes ON ttypes.ticket_type_id = t.ticket_type_id
JOIN events e ON e.event_id = ttypes.event_id
JOIN ticket_orders tord ON tord.order_id = items.order_id
JOIN attendees att ON att.attendee_id = tord.attendee_id
JOIN users u ON u.user_id = att.user_id
WHERE items.order_id = $1;

-- name: UpdateAdmitTicketOrder :exec
UPDATE ticket_orders
SET admit_status = true
WHERE order_id=$1;
