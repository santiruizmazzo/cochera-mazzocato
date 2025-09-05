package integration

import (
	"cochera/domain/tenant"
	"cochera/infrastructure"
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

	repo := infrastructure.NewPostgresTenantRepository(db)

	localTenant := tenant.NewTenantBuilder().WithID(1).Build()

	savedTenant, err := repo.Save(localTenant)
	if err != nil {
		t.Fatal(err)
	}

	if *localTenant != *savedTenant {
		t.Fatal("Expected tenant is different from saved tenant")
	}
}

func TestPostgresTenantRepository_Save_Fails_DuplicateDNI_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	existingTenant := tenant.NewTenantBuilder().WithEmail("a@b.com").Build()

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := tenant.NewTenantBuilder().WithEmail("c@d.com").Build()

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

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	existingTenant := tenant.NewTenantBuilder().WithDNI(22222222).Build()

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

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	existingTenant := tenant.NewTenantBuilder().WithDNI(1).Build()

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := tenant.NewTenantBuilder().WithDNI(2).Build()

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

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	existingTenant := tenant.NewTenantBuilder().WithEmail("neo@cortex.com").Build()

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

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	tenant, err := repo.GetTenantByID(666)
	if err == nil {
		t.Fatal("GetTenantByID should return error when it tenant does not exist")
	}

	if tenant != nil {
		t.Fatal("Tenant shouldn't be created when it is not found")
	}
}

func TestPostgresTenantRepository_GetAllTenants_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infrastructure.NewPostgresTenantRepository(db)

	existingTenant1 := tenant.NewTenantBuilder().WithDNI(433).WithEmail("first@tenant.com").Build()
	_, err := repo.Save(existingTenant1)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant2 := tenant.NewTenantBuilder().WithDNI(442).WithEmail("second@tenant.com").Build()
	_, err = repo.Save(existingTenant2)
	if err != nil {
		t.Fatal(err)
	}

	tenants, err := repo.GetAllTenants()
	if err != nil {
		t.Fatal("GetAllTenants shouldn't fail when there exists tenants: ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Returned a tenants list with a different size as expected")
	}

	if *existingTenant1 != *tenants[0] {
		t.Fatalf("Got %v, expected %v", tenants[0], existingTenant1)
	}

	if *existingTenant2 != *tenants[1] {
		t.Fatalf("Got %v, expected %v", tenants[1], existingTenant2)
	}
}
