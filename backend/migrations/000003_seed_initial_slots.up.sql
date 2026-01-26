BEGIN;

INSERT INTO slots (slot_number, tenant_id)
SELECT generate_series(1, 12), NULL
ON CONFLICT (slot_number) DO NOTHING;

COMMIT;
