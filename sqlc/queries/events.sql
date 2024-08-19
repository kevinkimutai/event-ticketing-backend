-- name: GetEvent :one
SELECT * FROM events
WHERE event_id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: ListUpcomingEvents :many
SELECT * 
FROM events
WHERE (date > NOW() OR (date = NOW() AND to_time > NOW()))
  AND (name ILIKE '%' || $5 || '%' OR $5 IS NULL)
  AND (category_id = $3 OR $3 IS NULL)
  AND (location_id = $4 OR $4 IS NULL)
ORDER BY date, to_time
LIMIT $1 OFFSET $2;


-- name: GetTotalEventsCount :one
SELECT COUNT(*)
FROM events
WHERE (date > NOW() OR (date = NOW()::DATE AND to_time > NOW()))
  AND (name ILIKE '%' || $3 || '%' OR $3 IS NULL)
  AND (category_id = $1 OR $1 IS NULL)
  AND (location_id = $2 OR $2 IS NULL);

-- name: CreateEvent :one
INSERT INTO events (
  name,category_id,date,from_time,to_time,location,description,longitude,latitude,poster_url,location_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9 ,$10,$11
)
RETURNING *;

-- name: UpdateEvent :exec
UPDATE events 
  set name = $2,
  date = $3,
  from_time = $4,
  to_time = $5,
  location = $6,
  description = $7,
  longitude = $8,
  latitude = $9,
  poster_url =$10,
  location_id = $11
WHERE event_id = $1;

-- name: DeleteCompany :exec
DELETE FROM events
WHERE event_id = $1;