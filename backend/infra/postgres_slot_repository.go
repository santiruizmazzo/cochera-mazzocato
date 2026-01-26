package infra

import (
	ent "cochera/domain/entities"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSlotRepository struct {
	db *pgxpool.Pool
}

func NewPostgresSlotRepository(db *pgxpool.Pool) *PostgresSlotRepository {
	return &PostgresSlotRepository{db: db}
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

	var id, number, tenantID int

	if err := row.Scan(&id, &number, &tenantID); err != nil {
		return nil, err
	}

	return ent.NewSlot(id, number, tenantID)
}
