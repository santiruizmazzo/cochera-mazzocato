package postgres

import (
	"cochera/internal/domain/calendar"
	"cochera/internal/domain/tenant"
	myerrors "cochera/internal/errors"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTenantRepository(db *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (repo *PostgresTenantRepository) GetTenantByID(id int) (*tenant.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE id = $1;`

	row := repo.db.QueryRow(context.Background(), query, id)

	return createTenantFromRow(row)
}

func createTenantFromRow(row pgx.Row) (*tenant.Tenant, error) {
	var tenant tenant.Tenant
	var entryMonth string
	var address, phone, email *string

	err := row.Scan(&tenant.ID, &tenant.DNI, &tenant.Name, &tenant.LastName, &address, &phone, &email, &entryMonth)
	if err != nil {
		return nil, myerrors.ErrTenantNotFound
	}

	tenant.Address = pointerToString(address)
	tenant.Phone = pointerToString(phone)
	tenant.Email = pointerToString(email)

	tenant.EntryMonth, err = calendar.NewMonthOfYearFromString(entryMonth)
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

func pointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (repo *PostgresTenantRepository) GetAllTenants() ([]*tenant.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants;`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*tenant.Tenant, error) {
		return createTenantFromRow(row)
	})
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
	if email == "" {
		return false, nil
	}

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

	var storableAddress any
	if tenant.Address == "" {
		storableAddress = nil
	} else {
		storableAddress = tenant.Address
	}

	var storablePhone any
	if tenant.Phone == "" {
		storablePhone = nil
	} else {
		storablePhone = tenant.Phone
	}

	var storableEmail any
	if tenant.Email == "" {
		storableEmail = nil
	} else {
		storableEmail = tenant.Email
	}

	row := repo.db.QueryRow(context.Background(), query, tenant.DNI, tenant.Name, tenant.LastName, storableAddress, storablePhone, storableEmail, tenant.EntryMonth.String())

	if err := row.Scan(&tenant.ID); err != nil {
		return nil, err
	}
	return tenant, nil
}
