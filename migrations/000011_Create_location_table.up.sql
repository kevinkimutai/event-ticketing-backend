-- Create the location table
CREATE TABLE location (
    location_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Add location_id column to events table
ALTER TABLE events 
ADD COLUMN location_id BIGINT NOT NULL;

-- Add foreign key constraint
ALTER TABLE events
ADD CONSTRAINT fk_location
FOREIGN KEY (location_id) 
REFERENCES location(location_id);
