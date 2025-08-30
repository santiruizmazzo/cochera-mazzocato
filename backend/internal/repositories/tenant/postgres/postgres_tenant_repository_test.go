package postgres

import (
	"cochera/internal/domain/tenant"
	"cochera/internal/domain/time"
	"cochera/tests/utils"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool
var err error

func TestMain(m *testing.M) {
	code := 1

	defer func() {
		os.Exit(code)
	}()

	db, err = utils.SetupTestDatabase()
	if err != nil {
		log.Printf("Failed connecting to test database: %v", err)
		return
	}
	defer utils.CleanupAndCloseTestDatabase(db)

	code = m.Run()
}

func TestPostgresTenantRepository_Save_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	localTenant := &tenant.Tenant{
		DNI:        12345678,
		Name:       "Manolo",
		LastName:   "Lamas",
		Address:    "Avenida Siempreviva 555",
		Phone:      "+5645551114",
		Email:      "mlamas@fifa09.com",
		EntryMonth: time.NewMonthOfYear(8, 2025),
	}

	savedTenant, err := repo.Save(localTenant)
	if err != nil {
		t.Fatal(err)
	}

	localTenant.ID = 1
	if savedTenant != localTenant {
		t.Fatal("Expected tenant is different from saved tenant")
	}
}

func TestPostgresTenantRepository_Save_Fails_DuplicateDNI_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	existingTenant := &tenant.Tenant{
		DNI:        33333333,
		Name:       "Manolo",
		LastName:   "Lamas",
		Address:    "Avenida Siempreviva 555",
		Phone:      "+5645551114",
		Email:      "manololamas@gmail.com",
		EntryMonth: time.NewMonthOfYear(8, 2025),
	}

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := &tenant.Tenant{
		DNI:        33333333,
		Name:       "Solid",
		LastName:   "Snake",
		Address:    "123 Stealth Mode St.",
		Phone:      "+5644440004",
		Email:      "solid@snake.com",
		EntryMonth: time.NewMonthOfYear(1, 2021),
	}

	createdTenant, err := repo.Save(newTenant)
	if err == nil {
		t.Fatal("Save should return error when it already exists a tenant with same DNI")
	}

	if createdTenant != nil {
		t.Fatal("Save should not return tenant when it already exists one with same DNI")
	}
}

func TestPostgresTenantRepository_ExistsTenantWithDNI_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	existingTenant := &tenant.Tenant{
		DNI:        22222222,
		Name:       "Manolo",
		LastName:   "Lamas",
		Address:    "Avenida Siempreviva 555",
		Phone:      "+5645551114",
		Email:      "mlamas@gmail.com",
		EntryMonth: time.NewMonthOfYear(8, 2025),
	}

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	dniAlreadyInUse, err := repo.ExistsTenantWithDNI(22222222)
	if err != nil || !dniAlreadyInUse {
		t.Fatal("Method should return that DNI already exists")
	}
}
