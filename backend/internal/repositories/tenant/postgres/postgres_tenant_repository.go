package postgres

import (
	"cochera/internal/domain/tenant"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTenantRepository(db *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (repo *PostgresTenantRepository) Save(tenant *tenant.Tenant) (*tenant.Tenant, error) {
	query := `
		INSERT INTO tenants (dni, name, last_name, address, phone, email, entry_month)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := repo.db.QueryRow(context.Background(), query, tenant.DNI, tenant.Name, tenant.LastName, tenant.Address, tenant.Phone, tenant.Email, tenant.EntryMonth.String())

	if err := row.Scan(&tenant.ID); err != nil {
		return nil, fmt.Errorf("failed inserting new tenant: %w", err)
	}
	return tenant, nil
}
