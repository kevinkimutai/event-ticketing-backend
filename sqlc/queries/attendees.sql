
-- name: CreateAttendee :one
INSERT INTO attendees (
  user_id
) VALUES (
  $1
)
RETURNING *;

-- name: GetAttendeeByUserID :one
SELECT * FROM attendees a
JOIN users u
ON a.user_id = u.user_id
WHERE a.attendee_id=$1
LIMIT 1; 

-- name: GetAttendee :one
SELECT * FROM attendees
WHERE attendee_id =$1
LIMIT 1;

-- name: GetAttendeeEvents :many
SELECT 
	att.attendee_id AS attendee_id,
	e.name AS event_name,
	e.date AS event_date,
	ord.payment_id AS payment_id,
	oitems.quantity AS quantity,
	oitems.total_price AS total_amount
  
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
JOIN 
    events e ON e.event_id = ttypes.event_id
JOIN 
    ticket_orders ord ON ord.order_id = oitems.order_id
JOIN 
    attendees att ON att.attendee_id = ord.attendee_id
JOIN 
    users u ON u.user_id = att.user_id
WHERE 
    att.user_id = $1
ORDER BY att.attendee_id DESC
LIMIT $2 OFFSET $3;

-- name: GetCountAttendeeEvents :one
SELECT 
	COUNT(*)
	   
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
JOIN 
    events e ON e.event_id = ttypes.event_id
JOIN 
    ticket_orders ord ON ord.order_id = oitems.order_id
JOIN 
    attendees att ON att.attendee_id = ord.attendee_id
JOIN 
    users u ON u.user_id = att.user_id
WHERE 
    att.user_id = $1;


-- name: GetEventsAttended :one
SELECT 
	COUNT(DISTINCT(e.event_id)) AS events_attended,
    SUM(oitems.total_price) AS total_spent
    
FROM 
    tickets t
JOIN 
    ticket_types ttypes ON t.ticket_type_id = ttypes.ticket_type_id
JOIN 
    ticket_order_items oitems ON oitems.ticket_id = t.ticket_id
JOIN 
    events e ON e.event_id = ttypes.event_id
JOIN 
    ticket_orders ord ON ord.order_id = oitems.order_id
JOIN 
    attendees att ON att.attendee_id = ord.attendee_id
JOIN 
    users u ON u.user_id = att.user_id
WHERE 
    att.user_id = $1;



