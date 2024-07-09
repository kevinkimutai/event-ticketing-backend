-- Make event_id column in ticket_types table NOT NULL
ALTER TABLE ticket_types
ALTER COLUMN event_id SET NOT NULL;
