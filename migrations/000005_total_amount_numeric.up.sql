ALTER TABLE ticket_orders
ALTER COLUMN total_amount TYPE numeric;

ALTER TABLE events
ADD COLUMN poster_url VARCHAR(255) NOT NULL;