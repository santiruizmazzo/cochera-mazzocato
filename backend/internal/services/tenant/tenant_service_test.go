package tenantservice

import (
	"cochera/internal/domain/tenant"
	myerrors "cochera/internal/errors"
	"encoding/json"
	"testing"
)

type mockTenantRepository struct {
	tenants map[int]*tenant.Tenant
	err     error
}

func (mockRepo *mockTenantRepository) GetAllTenants() ([]*tenant.Tenant, error) {
	var list []*tenant.Tenant
	for _, v := range mockRepo.tenants {
		list = append(list, v)
	}
	return list, nil
}

func (mockRepo *mockTenantRepository) GetTenantByID(id int) (*tenant.Tenant, error) {
	return nil, myerrors.ErrTenantNotFound
}

func (mockRepo *mockTenantRepository) Save(tenant *tenant.Tenant) (*tenant.Tenant, error) {
	if mockRepo.err != nil {
		return nil, mockRepo.err
	}
	tenant.ID = 1
	return tenant, nil
}

func (mockRepo *mockTenantRepository) ExistsTenantWithDNI(dni uint32) (bool, error) {
	for _, tenant := range mockRepo.tenants {
		if tenant != nil && tenant.DNI == dni {
			return true, nil
		}
	}
	return false, nil
}

func (mockRepo *mockTenantRepository) ExistsTenantWithEmail(email string) (bool, error) {
	for _, tenant := range mockRepo.tenants {
		if tenant != nil && tenant.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func TestTenantService_CreateTenant_Successfully(t *testing.T) {
	mockRepo := &mockTenantRepository{tenants: map[int]*tenant.Tenant{}}

	expectedTenant := tenant.NewTenantBuilder().WithID(1).Build()
	jsonTenant, _ := json.Marshal(expectedTenant)

	service := NewTenantService(mockRepo)

	tenant, err := service.CreateTenant(jsonTenant)
	if err != nil {
		t.Fatal(err)
	}

	if *expectedTenant != *tenant {
		t.Fatal("Expected tenant is different from created tenant")
	}
}

func TestTenantService_CreateTenant_Fails_DNIAlreadyExists(t *testing.T) {
	existingTenant := tenant.NewTenantBuilder().Build()
	mockRepo := &mockTenantRepository{tenants: map[int]*tenant.Tenant{1: existingTenant}}

	newTenant := tenant.NewTenantBuilder().WithEmail("another@email.com").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	service := NewTenantService(mockRepo)
	tenant, err := service.CreateTenant(jsonTenant)

	if tenant != nil {
		t.Fatal("Tenant should not be created")
	}

	if err != myerrors.ErrDuplicateDNI {
		t.Fatal("Error should be of type duplicate DNI")
	}
}

func TestTenantService_CreateTenant_Fails_NonNumericDNI(t *testing.T) {
	service := NewTenantService(&mockTenantRepository{})

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

	if err != myerrors.ErrDNIMustBeNumber {
		t.Fatal("Error should be of type DNI must be a number")
	}
}

func TestTenantService_GetTenantByID_Fails_TenantDoesNotExist(t *testing.T) {
	service := NewTenantService(&mockTenantRepository{})

	tenant, err := service.GetTenantByID(9)

	if tenant != nil {
		t.Fatal("Tenant should not be found")
	}

	if err != myerrors.ErrTenantNotFound {
		t.Fatal("Error should be of type tenant not found")
	}
}

func TestTenantService_GetTenants_Successfully(t *testing.T) {
	expectedTenant1 := tenant.NewTenantBuilder().WithID(1).WithDNI(43295798).WithEmail("1@2.com").Build()
	expectedTenant2 := tenant.NewTenantBuilder().WithID(2).WithDNI(41630284).WithEmail("3@4.com").Build()

	mockRepo := &mockTenantRepository{tenants: map[int]*tenant.Tenant{
		1: expectedTenant1,
		2: expectedTenant2,
	}}

	service := NewTenantService(mockRepo)

	tenants, err := service.GetAllTenants()

	if err != nil {
		t.Fatal("GetTenants shouldn't fail here")
	}

	if len(tenants) != 2 {
		t.Fatal("Incorrect number of tenants retrieved")
	}

	if *expectedTenant1 != *tenants[0] {
		t.Fatalf("Expected %v, got %v", expectedTenant1, tenants[0])
	}

	if *expectedTenant2 != *tenants[1] {
		t.Fatalf("Expected %v, got %v", expectedTenant2, tenants[1])
	}
}
