package postgres

import (
	"cochera/internal/domain/tenant"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTenantRepository(db *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (repo *PostgresTenantRepository) ExistsTenantWithDNI(dni uint32) (bool, error) {
	query := `SELECT COUNT(*) > 0 AS exists FROM tenants WHERE dni = $1;`

	var exists bool
	err := repo.db.QueryRow(context.Background(), query, dni).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (repo *PostgresTenantRepository) Save(tenant *tenant.Tenant) (*tenant.Tenant, error) {
	query := `
		INSERT INTO tenants (dni, name, last_name, address, phone, email, entry_month)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := repo.db.QueryRow(context.Background(), query, tenant.DNI, tenant.Name, tenant.LastName, tenant.Address, tenant.Phone, tenant.Email, tenant.EntryMonth.String())

	if err := row.Scan(&tenant.ID); err != nil {
		return nil, translateError(err)
	}
	return tenant, nil
}

func translateError(err error) error {
	errorMessage := err.Error()

	if strings.HasPrefix(errorMessage, "ERROR: duplicate key value violates unique constraint") {
		if strings.Contains(errorMessage, "tenants_dni_key") {
			return errors.New("dni already exists")
		}
		if strings.Contains(errorMessage, "tenants_email_key") {
			return errors.New("email already exists")
		}
	}

	return err
}
