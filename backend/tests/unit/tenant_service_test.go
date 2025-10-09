package unit

import (
	"cochera/application/services"
	"cochera/domain"
	"cochera/infra"
	"encoding/json"
	"testing"
)

type mockTenantRepository struct {
	tenants map[int]*domain.Tenant
	err     error
}

func (mockRepo *mockTenantRepository) GetAllTenants() ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range mockRepo.tenants {
		list = append(list, tenant)
	}
	return list, nil
}

func (mockRepo *mockTenantRepository) GetAllTenantsByName(name string) ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range mockRepo.tenants {
		if tenant.HasName(name) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (mockRepo *mockTenantRepository) GetAllTenantsByLastName(lastName string) ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range mockRepo.tenants {
		if tenant.HasLastName(lastName) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (mockRepo *mockTenantRepository) GetTenantByID(id int) (*domain.Tenant, error) {
	return nil, infra.ErrTenantNotFound
}

func (mockRepo *mockTenantRepository) Save(tenant *domain.Tenant) (*domain.Tenant, error) {
	if mockRepo.err != nil {
		return nil, mockRepo.err
	}
	tenant.ID = 1
	return tenant, nil
}

func (mockRepo *mockTenantRepository) ExistsTenantWithDNI(dni uint32) (bool, error) {
	for _, tenant := range mockRepo.tenants {
		if tenant != nil && tenant.HasDNI(dni) {
			return true, nil
		}
	}
	return false, nil
}

func (mockRepo *mockTenantRepository) ExistsTenantWithEmail(email string) (bool, error) {
	for _, tenant := range mockRepo.tenants {
		if tenant != nil && tenant.HasEmail(email) {
			return true, nil
		}
	}
	return false, nil
}

func TestTenantService_CreateTenant_Successfully(t *testing.T) {
	mockRepo := &mockTenantRepository{tenants: map[int]*domain.Tenant{}}

	expectedTenant := domain.NewTenantBuilder().WithID(1).Build()
	jsonTenant, _ := json.Marshal(expectedTenant)

	service := services.NewTenantService(mockRepo)

	tenant, err := service.CreateTenant(jsonTenant)
	if err != nil {
		t.Fatal(err)
	}

	if *expectedTenant != *tenant {
		t.Fatal("Expected tenant is different from created tenant")
	}
}

func TestTenantService_CreateTenant_Fails_DNIAlreadyExists(t *testing.T) {
	existingTenant := domain.NewTenantBuilder().Build()
	mockRepo := &mockTenantRepository{tenants: map[int]*domain.Tenant{1: existingTenant}}

	newTenant := domain.NewTenantBuilder().WithEmail("another@email.com").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	service := services.NewTenantService(mockRepo)
	tenant, err := service.CreateTenant(jsonTenant)

	if tenant != nil {
		t.Fatal("Tenant should not be created")
	}

	if err != services.ErrDuplicateDNI {
		t.Fatal("Error should be of type duplicate DNI")
	}
}

func TestTenantService_CreateTenant_Fails_NonNumericDNI(t *testing.T) {
	service := services.NewTenantService(&mockTenantRepository{})

	jsonTenant, _ := json.Marshal(map[string]any{
		"dni":         "hola",
		"name":        "Samwise",
		"last_name":   "Gamyi",
		"entry_month": "06-2025",
	})

	tenant, err := service.CreateTenant(jsonTenant)

	if tenant != nil {
		t.Fatal("Tenant should not be created")
	}

	if err != domain.ErrDNIMustBeANumber {
		t.Fatal("Error should be of type DNI must be a number", err)
	}
}

func TestTenantService_GetTenantByID_Fails_TenantDoesNotExist(t *testing.T) {
	service := services.NewTenantService(&mockTenantRepository{})

	tenant, err := service.GetTenantByID(9)

	if tenant != nil {
		t.Fatal("Tenant should not be found")
	}

	if err != infra.ErrTenantNotFound {
		t.Fatal("Error should be of type tenant not found")
	}
}

func TestTenantService_GetAllTenants_Successfully(t *testing.T) {
	expectedTenant1 := domain.NewTenantBuilder().WithID(1).WithDNI(43295798).WithEmail("1@2.com").Build()
	expectedTenant2 := domain.NewTenantBuilder().WithID(2).WithDNI(41630284).WithEmail("3@4.com").Build()

	mockRepo := &mockTenantRepository{tenants: map[int]*domain.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAllTenants()

	if err != nil {
		t.Fatal("GetTenants shouldn't fail here")
	}

	if len(tenants) != 2 {
		t.Fatal("Incorrect number of tenants retrieved")
	}
}

func TestTenantService_GetAllTenantsByName_Successfully(t *testing.T) {
	nameToFilter := "Salvatore"

	expectedTenant1 := domain.NewTenantBuilder().WithName(nameToFilter).Build()
	expectedTenant2 := domain.NewTenantBuilder().WithName("Giuseppe").Build()
	expectedTenant3 := domain.NewTenantBuilder().WithName(nameToFilter).Build()
	mockRepo := &mockTenantRepository{tenants: map[int]*domain.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
		3: expectedTenant3,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAllTenantsByName(nameToFilter)

	if err != nil {
		t.Fatal("GetAllTenantsByName should not fail: ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	if !tenants[0].HasName(nameToFilter) || !tenants[1].HasName(nameToFilter) {
		t.Fatal("Failed to get all tenants with same name")
	}
}

func TestTenantService_GetAllTenantsByLastName_Successfully(t *testing.T) {
	lastNameToFilter := "Leone"

	expectedTenant1 := domain.NewTenantBuilder().WithLastName(lastNameToFilter).Build()
	expectedTenant2 := domain.NewTenantBuilder().WithLastName("Cipriani").Build()
	expectedTenant3 := domain.NewTenantBuilder().WithLastName(lastNameToFilter).Build()
	mockRepo := &mockTenantRepository{tenants: map[int]*domain.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
		3: expectedTenant3,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAllTenantsByLastName(lastNameToFilter)

	if err != nil {
		t.Fatal("GetAllTenantsByLastName should not fail: ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	if !tenants[0].HasLastName(lastNameToFilter) || !tenants[1].HasLastName(lastNameToFilter) {
		t.Fatal("Failed to get all tenants with same last name")
	}
}
