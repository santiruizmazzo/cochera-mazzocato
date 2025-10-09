package infra

import (
	ent "cochera/domain/entities"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

var (
	ErrTenantNotFound         = errors.New("inquilino no encontrado")
	ErrNoMatchingTenantsFound = errors.New("no se encontraron inquilinos que coincidan")
)

func NewPostgresTenantRepository(db *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (repo PostgresTenantRepository) GetByID(id int) (*ent.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE id = $1;`

	row := repo.db.QueryRow(context.Background(), query, id)

	return repo.createTenantFromRow(row)
}

func (repo PostgresTenantRepository) createTenantFromRow(row pgx.Row) (*ent.Tenant, error) {
	var (
		id, dni                    int
		name, lastName, entryMonth string
		address, phone, email      *string
	)

	err := row.Scan(&id, &dni, &name, &lastName, &address, &phone, &email, &entryMonth)
	if err != nil {
		return nil, ErrTenantNotFound
	}

	return ent.NewTenant(id, dni, name, lastName, address, phone, email, entryMonth)
}

func (repo PostgresTenantRepository) GetAll() ([]*ent.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants;`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repo.createListOfTenantsFromRows(rows)
}

func (repo PostgresTenantRepository) GetAllWithName(name string) ([]*ent.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE name ILIKE $1;`

	wildcardString := "%" + name + "%"
	rows, err := repo.db.Query(context.Background(), query, wildcardString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repo.createListOfTenantsFromRows(rows)
}

func (repo PostgresTenantRepository) GetAllWithLastName(lastName string) ([]*ent.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE last_name ILIKE $1;`

	wildcardString := "%" + lastName + "%"
	rows, err := repo.db.Query(context.Background(), query, wildcardString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repo.createListOfTenantsFromRows(rows)
}

func (repo PostgresTenantRepository) createListOfTenantsFromRows(rows pgx.Rows) ([]*ent.Tenant, error) {
	tenants, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*ent.Tenant, error) {
		return repo.createTenantFromRow(row)
	})
	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, ErrNoMatchingTenantsFound
	}

	return tenants, nil
}

func (repo PostgresTenantRepository) ExistsTenantWithDNI(dni uint32) (bool, error) {
	query := `SELECT COUNT(*) > 0 AS exists FROM tenants WHERE dni = $1;`

	var exists bool
	err := repo.db.QueryRow(context.Background(), query, dni).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (repo PostgresTenantRepository) ExistsTenantWithEmail(email string) (bool, error) {
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

func (repo PostgresTenantRepository) Save(tenant *ent.Tenant) (*ent.Tenant, error) {
	query := `
		INSERT INTO tenants (dni, name, last_name, address, phone, email, entry_month)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := repo.db.QueryRow(context.Background(), query, tenant.GetDNI(), tenant.GetName(), tenant.GetLastName(), tenant.GetAddress(), tenant.GetPhone(), tenant.GetEmail(), tenant.GetEntryMonth())

	var id int

	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	tenant.SetID(id)
	return tenant, nil
}
