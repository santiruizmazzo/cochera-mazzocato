package postgres

import (
	"cochera/internal/domain/calendar"
	"cochera/internal/domain/tenant"
	myerrors "cochera/internal/errors"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTenantRepository(db *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (repo *PostgresTenantRepository) GetTenantByID(id int) (*tenant.Tenant, error) {
	query := `SELECT * FROM tenants WHERE id = $1;`

	row := repo.db.QueryRow(context.Background(), query, id)

	var tenant tenant.Tenant
	var rawEntryMonth string

	err := row.Scan(&tenant.ID, &tenant.DNI, &tenant.Name, &tenant.LastName, &tenant.Address, &tenant.Phone, &tenant.Email, &rawEntryMonth)
	if err != nil {
		return nil, myerrors.ErrTenantNotFound
	}

	tenant.EntryMonth, err = calendar.NewMonthOfYearFromString(rawEntryMonth)
	if err != nil {
		return nil, err
	}

	return &tenant, nil
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

func (repo *PostgresTenantRepository) ExistsTenantWithEmail(email string) (bool, error) {
	query := `SELECT COUNT(*) > 0 AS exists FROM tenants WHERE email = $1;`

	var exists bool
	err := repo.db.QueryRow(context.Background(), query, email).Scan(&exists)
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
		return nil, err
	}
	return tenant, nil
}
