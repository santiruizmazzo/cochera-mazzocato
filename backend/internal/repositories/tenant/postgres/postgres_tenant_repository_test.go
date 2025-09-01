package postgres

import (
	"cochera/internal/domain/calendar"
	"cochera/internal/domain/tenant"
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
		EntryMonth: calendar.NewMonthOfYear(8, 2025),
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
		EntryMonth: calendar.NewMonthOfYear(8, 2025),
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
		EntryMonth: calendar.NewMonthOfYear(1, 2021),
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
		EntryMonth: calendar.NewMonthOfYear(8, 2025),
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

func TestPostgresTenantRepository_Save_Fails_DuplicateEmail_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	existingTenant := &tenant.Tenant{
		DNI:        123,
		Name:       "Max",
		LastName:   "Payne",
		Address:    "745 Suicidal Av.",
		Phone:      "+56423456714",
		Email:      "max@payne.com",
		EntryMonth: calendar.NewMonthOfYear(6, 2025),
	}

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := &tenant.Tenant{
		DNI:        4231,
		Name:       "Mack",
		LastName:   "Paind",
		Address:    "123 Full Gaga St.",
		Phone:      "+56111111",
		Email:      "max@payne.com",
		EntryMonth: calendar.NewMonthOfYear(1, 2021),
	}

	createdTenant, err := repo.Save(newTenant)
	if err == nil {
		t.Fatal("Save should return error when it already exists a tenant with same email")
	}

	if createdTenant != nil {
		t.Fatal("Save should not return tenant when it already exists one with same email")
	}
}

func TestPostgresTenantRepository_ExistsTenantWithEmail_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	existingTenant := &tenant.Tenant{
		DNI:        777777777,
		Name:       "Neo",
		LastName:   "Cortex",
		Address:    "123 I hate Crash st.",
		Phone:      "+1999999999",
		Email:      "neo@cortex.com",
		EntryMonth: calendar.NewMonthOfYear(8, 2025),
	}

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	emailAlreadyInUse, err := repo.ExistsTenantWithEmail("neo@cortex.com")
	if err != nil || !emailAlreadyInUse {
		t.Fatal("Method should return that email is already in use")
	}
}

func TestPostgresTenantRepository_GetTenantByID_Fails_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repo := NewPostgresTenantRepository(db)

	tenant, err := repo.GetTenantByID(666)
	if err == nil {
		t.Fatal("GetTenantByID should return error when it tenant does not exist")
	}

	if tenant != nil {
		t.Fatal("Tenant shouldn't be created when it is not found")
	}
}
