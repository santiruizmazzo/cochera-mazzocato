package infra

import (
	ent "cochera/domain/entities"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSlotRepository struct {
	db *pgxpool.Pool
}

var (
	ErrSlotNotFound = errors.New("plaza no encontrada")
	// ErrNoMatchingSlotsFound = errors.New("no se encontraron plazas que coincidan")
)

func NewPostgresSlotRepository(db *pgxpool.Pool) *PostgresSlotRepository {
	return &PostgresSlotRepository{db: db}
}

func (repo PostgresSlotRepository) GetByID(id int) (*ent.Slot, error) {
	query := `SELECT id, slot_number, tenant_id FROM slots WHERE id = $1;`

	row := repo.db.QueryRow(context.Background(), query, id)

	return repo.createSlotFromRow(row)
}

func (repo PostgresSlotRepository) createSlotFromRow(row pgx.Row) (*ent.Slot, error) {
	var (
		id, slotNumber   int
		nullableTenantID *int
	)

	err := row.Scan(&id, &slotNumber, &nullableTenantID)
	if err != nil {
		return nil, ErrSlotNotFound
	}

	var tenantID int
	if nullableTenantID != nil {
		tenantID = *nullableTenantID
	}

	return ent.NewSlot(id, slotNumber, tenantID)
}

func (repo PostgresSlotRepository) Save(slot *ent.Slot) (*ent.Slot, error) {
	query := `
		INSERT INTO slots (slot_number, tenant_id)
		VALUES ($1, $2)
		ON CONFLICT (slot_number) 
		DO UPDATE SET 
			tenant_id = EXCLUDED.tenant_id,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id, slot_number, tenant_id
	`

	row := repo.db.QueryRow(context.Background(), query, slot.Number, slot.TenantID)

	return repo.createSlotFromRow(row)
}
