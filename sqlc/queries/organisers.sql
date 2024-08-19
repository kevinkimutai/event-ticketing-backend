
-- name: CreateOrganiser :one
INSERT INTO organisers (
  user_id,event_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetOrganisersByUserID :many
SELECT * FROM organisers 
WHERE user_id = $1;

-- name: GetCountOrganisersByUserID :many
SELECT COUNT(*) FROM organisers
WHERE user_id = $1;

-- name: GetOrganisersEventByID :many
SELECT 
  att.attendee_id AS attendee_id,
	u.full_name AS fullname,
	u.email AS email,
	ttypes.name AS ticket_type_name,
	oitems.quantity AS quantity,
	oitems.total_price AS total
	
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
JOIN 
	ticket_orders ord ON ord.order_id = oitems.order_id
JOIN 
	attendees att ON att.attendee_id = ord.attendee_id
JOIN 
	users u ON u.user_id = att.user_id

WHERE 
    ttypes.event_id = $1;


-- name: SumAmountEvents :one
SELECT 
    SUM(oitems.total_price) AS total_amount
FROM 
    tickets t
INNER JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
INNER JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
INNER JOIN 
    organisers org ON org.event_id = ttypes.event_id
WHERE 
    org.user_id = $1;


-- name: GetOrganisersEventSums :one
SELECT 
    COALESCE(SUM(oitems.quantity), 0) AS total_sold_tickets,
    COALESCE(SUM(oitems.total_price), 0) AS total_price
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
WHERE 
    ttypes.event_id = $1;


-- name: GetOrganisersEventCount :one
SELECT 
 COUNT(*)
	
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
JOIN 
	ticket_orders ord ON ord.order_id = oitems.order_id
JOIN 
	attendees att ON att.attendee_id = ord.attendee_id
JOIN 
	users u ON u.user_id = att.user_id

WHERE 
    ttypes.event_id = $1;


