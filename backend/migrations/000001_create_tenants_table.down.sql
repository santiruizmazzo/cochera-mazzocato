BEGIN;

DROP TRIGGER IF EXISTS update_tenants_updated_at ON tenants;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_tenants_is_active;

DROP TABLE IF EXISTS tenants;

COMMIT;