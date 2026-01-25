package infra

import (
	ent "cochera/domain/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSlotRepository struct {
	db *pgxpool.Pool
}

func NewPostgresSlotRepository(db *pgxpool.Pool) *PostgresSlotRepository {
	return &PostgresSlotRepository{db: db}
}

func (repo PostgresSlotRepository) Save(slot *ent.Slot) (*ent.Slot, error) {
	panic("unimplemented")
}
