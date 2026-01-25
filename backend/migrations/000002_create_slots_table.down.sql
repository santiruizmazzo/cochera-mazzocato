BEGIN;

DROP TRIGGER IF EXISTS update_slots_updated_at ON slots;

DROP INDEX IF EXISTS idx_slots_tenant_id;
DROP INDEX IF EXISTS idx_slots_available;
DROP INDEX IF EXISTS idx_slots_slot_number;

DROP TABLE IF EXISTS slots;

COMMIT;