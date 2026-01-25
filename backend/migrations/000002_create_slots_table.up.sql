BEGIN;

CREATE TABLE IF NOT EXISTS slots (
    id SERIAL PRIMARY KEY,
    slot_number INTEGER NOT NULL UNIQUE,
    tenant_id INTEGER NOT NULL REFERENCES tenants(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_slots_tenant_id ON slots(tenant_id);
CREATE INDEX IF NOT EXISTS idx_slots_slot_number ON slots(slot_number);

-- Crear función solo si no existe (ya debería existir de la migración de tenants)
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_proc 
        WHERE proname = 'update_updated_at_column'
    ) THEN
        CREATE FUNCTION update_updated_at_column()
        RETURNS TRIGGER AS $func$
        BEGIN
            NEW.updated_at = CURRENT_TIMESTAMP;
            RETURN NEW;
        END;
        $func$ LANGUAGE 'plpgsql';
    END IF;
END $$;

-- Trigger para actualizar updated_at automáticamente
DROP TRIGGER IF EXISTS update_slots_updated_at ON slots;
CREATE TRIGGER update_slots_updated_at
    BEFORE UPDATE ON slots
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

COMMIT;