ALTER TABLE ticket_orders
DROP CONSTRAINT ticket_orders_user_id_fkey,  -- Drop the foreign key constraint on user_id
DROP COLUMN user_id,                         -- Drop the user_id column
ADD COLUMN attendee_id BIGINT,               -- Add the attendee_id column
ADD COLUMN total_amount FLOAT,               -- Add the total_amount column
ADD FOREIGN KEY (attendee_id) REFERENCES attendees(attendee_id) ON DELETE CASCADE;  -- Add the foreign key constraint for attendee_id
