ALTER TABLE IF EXISTS keepers
    DROP COLUMN IF EXISTS location_id,
    ADD location VARCHAR;