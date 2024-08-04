-- Remove foreign key constraint from the events table
ALTER TABLE events 
DROP CONSTRAINT IF EXISTS fk_location;

-- Remove the location_id column from the events table
ALTER TABLE events 
DROP COLUMN IF EXISTS location_id;

-- Drop the location table
DROP TABLE IF EXISTS location;
