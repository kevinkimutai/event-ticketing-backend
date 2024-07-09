ALTER TABLE tickets 
DROP COLUMN event_id;

ALTER TABLE ticket_types
ADD COLUMN event_id BIGINT,
ADD CONSTRAINT fk_event_id
    FOREIGN KEY (event_id)
    REFERENCES events(event_id)
    ON DELETE CASCADE;
