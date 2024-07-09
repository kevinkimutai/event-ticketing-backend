-- Step 1: Drop the existing foreign key constraint
ALTER TABLE ticket_orders
DROP CONSTRAINT ticket_orders_payment_id_fkey;

-- Step 2: Alter the column to allow NULL values
ALTER TABLE ticket_orders
ALTER COLUMN payment_id DROP NOT NULL;

-- Step 3: Add the new foreign key constraint with ON DELETE SET NULL action
ALTER TABLE ticket_orders
ADD CONSTRAINT ticket_orders_payment_id_fkey
FOREIGN KEY (payment_id) REFERENCES payments(payment_id) ON DELETE SET NULL;
