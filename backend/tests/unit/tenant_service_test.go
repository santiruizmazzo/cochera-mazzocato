package unit

import (
	"cochera/application/services"
	"cochera/domain"
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
	"cochera/infra"
	"encoding/json"
	"testing"
)

func TestTenantService_CreateTenant_Successfully(t *testing.T) {
	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{}}

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
	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{1: existingTenant}}

	newTenant := domain.NewTenantBuilder().WithEmail("another@email.com").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	service := services.NewTenantService(mockRepo)
	tenant, err := service.CreateTenant(jsonTenant)

	if tenant != nil {
		t.Fatal("Tenant should not be created")
	}

	if err != services.ErrDuplicateDNI {
		t.Fatal("Error should be of type duplicate DNI: ", err)
	}
}

func TestTenantService_CreateTenant_Fails_NonNumericDNI(t *testing.T) {
	service := services.NewTenantService(&infra.InMemoryTenantRepository{})

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

	if err != vo.ErrDNIMustBeAnInteger {
		t.Fatal("Error should be of type DNI must be a number", err)
	}
}

func TestTenantService_GetTenantByID_Fails_TenantDoesNotExist(t *testing.T) {
	service := services.NewTenantService(&infra.InMemoryTenantRepository{})

	tenant, err := service.GetByID(9)

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

	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAll()

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
	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
		3: expectedTenant3,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAllWithName(nameToFilter)

	if err != nil {
		t.Fatal("GetAllWithName should not fail: ", err)
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
	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
		3: expectedTenant3,
	}}

	service := services.NewTenantService(mockRepo)

	tenants, err := service.GetAllWithLastName(lastNameToFilter)

	if err != nil {
		t.Fatal("GetAllWithLastName should not fail: ", err)
	}

	if len(tenants) != 2 {
		t.Fatal("Expected tenants list of size 2, got ", len(tenants))
	}

	if !tenants[0].HasLastName(lastNameToFilter) || !tenants[1].HasLastName(lastNameToFilter) {
		t.Fatal("Failed to get all tenants with same last name")
	}
}

func TestTenantService_ModifyByID_Successfully(t *testing.T) {
	existingTenant := domain.NewTenantBuilder().Build()
	mockRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{1: existingTenant}}

	service := services.NewTenantService(mockRepo)

	requestBody, _ := json.Marshal(map[string]any{"dni": 666})

	modifiedTenant, err := service.ModifyByID(1, requestBody)

	if err != nil {
		t.Fatal("ModifyByID should not fail: ", err)
	}

	if modifiedTenant.DNI != 666 {
		t.Fatal("Expected new DNI value of 666, got", modifiedTenant.DNI)
	}
}
