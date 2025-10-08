package infra

import (
	"cochera/domain"
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

func (repo *PostgresTenantRepository) GetTenantByID(id int) (*domain.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE id = $1;`

	row := repo.db.QueryRow(context.Background(), query, id)

	return createTenantFromRow(row)
}

func createTenantFromRow(row pgx.Row) (*domain.Tenant, error) {
	// var tenant domain.Tenant
	// var entryMonth string
	// var address, phone, email *string

	var (
		id         int
		dni        int
		name       string
		lastName   string
		address    *string
		phone      *string
		email      *string
		entryMonth string
	)

	err := row.Scan(&id, &dni, &name, &lastName, &address, &phone, &email, &entryMonth)
	if err != nil {
		return nil, domain.ErrTenantNotFound
	}

	// var newAddress = pointerToString(address)
	// var newPhone = pointerToString(phone)
	// var newEmail = pointerToString(email)

	// tenant.EntryMonth, err = domain.NewMonthOfYearFromString(entryMonth)
	// if err != nil {
	// 	return nil, err
	// }
	return domain.NewTenant(id, dni, name, lastName, address, phone, email, entryMonth)
}

// func pointerToString(s *string) string {
// 	if s == nil {
// 		return ""
// 	}
// 	return *s
// }

func (repo *PostgresTenantRepository) GetAllTenants() ([]*domain.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants;`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return createListOfTenantsFromRows(rows)
}

func (repo *PostgresTenantRepository) GetAllTenantsByName(name string) ([]*domain.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE name ILIKE $1;`

	wildcardString := "%" + name + "%"
	rows, err := repo.db.Query(context.Background(), query, wildcardString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return createListOfTenantsFromRows(rows)
}

func (repo *PostgresTenantRepository) GetAllTenantsByLastName(lastName string) ([]*domain.Tenant, error) {
	query := `SELECT id, dni, name, last_name, address, phone, email, entry_month FROM tenants WHERE last_name ILIKE $1;`

	wildcardString := "%" + lastName + "%"
	rows, err := repo.db.Query(context.Background(), query, wildcardString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return createListOfTenantsFromRows(rows)
}

func createListOfTenantsFromRows(rows pgx.Rows) ([]*domain.Tenant, error) {
	tenants, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*domain.Tenant, error) {
		return createTenantFromRow(row)
	})
	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, domain.ErrNoMatchingTenantsFound
	}

	return tenants, nil
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

func (repo *PostgresTenantRepository) Save(tenant *domain.Tenant) (*domain.Tenant, error) {
	query := `
		INSERT INTO tenants (dni, name, last_name, address, phone, email, entry_month)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var storableAddress any
	if tenant.GetAddress() == "" {
		storableAddress = nil
	} else {
		storableAddress = tenant.GetAddress()
	}

	var storablePhone any
	if tenant.GetPhone() == "" {
		storablePhone = nil
	} else {
		storablePhone = tenant.GetPhone()
	}

	var storableEmail any
	if tenant.GetEmail() == "" {
		storableEmail = nil
	} else {
		storableEmail = tenant.GetEmail()
	}

	row := repo.db.QueryRow(context.Background(), query, tenant.GetDNI(), tenant.GetName(), tenant.GetLastName(), storableAddress, storablePhone, storableEmail, tenant.GetEntryMonth())

	var id int

	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	tenant.SetID(id)
	return tenant, nil
}
