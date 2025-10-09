package integration

import (
	"cochera/domain"
	"cochera/infra"
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

	repo := infra.NewPostgresTenantRepository(db)

	localTenant := domain.NewTenantBuilder().WithID(1).Build()

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
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant := domain.NewTenantBuilder().WithEmail("a@b.com").Build()

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := domain.NewTenantBuilder().WithEmail("c@d.com").Build()

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
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant := domain.NewTenantBuilder().WithDNI(22222222).Build()

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	dniAlreadyInUse, err := repo.ExistsWithDNI(22222222)
	if err != nil || !dniAlreadyInUse {
		t.Fatal("Method should return that DNI already exists")
	}
}

func TestPostgresTenantRepository_Save_Fails_DuplicateEmail_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant := domain.NewTenantBuilder().WithDNI(1).Build()

	_, err := repo.Save(existingTenant)
	if err != nil {
		t.Fatal(err)
	}

	newTenant := domain.NewTenantBuilder().WithDNI(2).Build()

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
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant := domain.NewTenantBuilder().WithEmail("neo@cortex.com").Build()

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
	repo := infra.NewPostgresTenantRepository(db)

	tenant, err := repo.GetByID(666)
	if err == nil {
		t.Fatal("GetByID should return error when it tenant does not exist")
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
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant1 := domain.NewTenantBuilder().WithDNI(433).WithEmail("first@tenant.com").Build()
	_, err := repo.Save(existingTenant1)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant2 := domain.NewTenantBuilder().WithDNI(442).WithEmail("second@tenant.com").Build()
	_, err = repo.Save(existingTenant2)
	if err != nil {
		t.Fatal(err)
	}

	tenants, err := repo.GetAll()
	if err != nil {
		t.Fatal("GetAll shouldn't fail when there exists tenants: ", err)
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

func TestPostgresTenantRepository_GetAllTenantsByName_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infra.NewPostgresTenantRepository(db)

	nameToFilter := "Toni"

	existingTenant1 := domain.NewTenantBuilder().WithName(nameToFilter).Build()
	_, err := repo.Save(existingTenant1)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant2 := domain.NewTenantBuilder().WithDNI(2).WithEmail("second@tenant.com").Build()
	_, err = repo.Save(existingTenant2)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant3 := domain.NewTenantBuilder().WithName(nameToFilter).WithDNI(3).WithEmail("third@tenant.com").Build()
	_, err = repo.Save(existingTenant3)
	if err != nil {
		t.Fatal(err)
	}

	tenants, err := repo.GetAllWithName(nameToFilter)
	if err != nil {
		t.Fatal("GetAllWithName shouldn't fail when there exists tenants with that name: ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	if !tenants[0].HasName(nameToFilter) || !tenants[1].HasName(nameToFilter) {
		t.Fatal("Names of filtered tenants are incorrect")
	}
}

func TestPostgresTenantRepository_GetAllTenantsByName_MatchPartially_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant1 := domain.NewTenantBuilder().WithName("Mario").Build()
	_, err := repo.Save(existingTenant1)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant2 := domain.NewTenantBuilder().WithDNI(2).WithEmail("second@tenant.com").Build()
	_, err = repo.Save(existingTenant2)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant3 := domain.NewTenantBuilder().WithName("Lamar").WithDNI(3).WithEmail("third@tenant.com").Build()
	_, err = repo.Save(existingTenant3)
	if err != nil {
		t.Fatal(err)
	}

	nameToFilter := "Mar"
	tenants, err := repo.GetAllWithName(nameToFilter)
	if err != nil {
		t.Fatal("GetAllWithName shouldn't fail when there exists tenants with a name that matches 'Mar': ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	utils.AssertResponseStringContains(tenants[0].GetName(), nameToFilter, t)
	utils.AssertResponseStringContains(tenants[1].GetName(), nameToFilter, t)
}

func TestPostgresTenantRepository_GetAllTenantsByLastName_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	utils.CleanupTestDatabase(db)
	repo := infra.NewPostgresTenantRepository(db)

	existingTenant1 := domain.NewTenantBuilder().Build()
	_, err := repo.Save(existingTenant1)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant2 := domain.NewTenantBuilder().WithDNI(2).WithEmail("second@tenant.com").Build()
	_, err = repo.Save(existingTenant2)
	if err != nil {
		t.Fatal(err)
	}

	existingTenant3 := domain.NewTenantBuilder().WithLastName("Jaoming").WithDNI(3).WithEmail("third@tenant.com").Build()
	_, err = repo.Save(existingTenant3)
	if err != nil {
		t.Fatal(err)
	}

	lastNameToFilter := "Lee"
	tenants, err := repo.GetAllWithLastName(lastNameToFilter)
	if err != nil {
		t.Fatal("GetAllWithLastName shouldn't fail when there exists tenants with last name 'Lee': ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	utils.AssertResponseStringContains(tenants[0].GetLastName(), lastNameToFilter, t)
	utils.AssertResponseStringContains(tenants[1].GetLastName(), lastNameToFilter, t)
}
