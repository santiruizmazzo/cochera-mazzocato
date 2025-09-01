package tenantservice

import (
	"cochera/internal/domain/calendar"
	"cochera/internal/domain/tenant"
	myerrors "cochera/internal/errors"
	"testing"
)

type mockTenantRepository struct {
	tenants map[int]*tenant.Tenant
	err     error
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

	expectedTenant := &tenant.Tenant{
		ID:         1,
		DNI:        12345678,
		Name:       "Manolo",
		LastName:   "Lamas",
		Address:    "Avenida Siempreviva 555",
		Phone:      "+5645551114",
		Email:      "mlamas@fifa09.com",
		EntryMonth: calendar.NewMonthOfYear(8, 2025),
	}
	jsonTenant := []byte(`{"dni":12345678,"name":"Manolo","last_name":"Lamas","address":"Avenida Siempreviva 555","phone":"+5645551114","email":"mlamas@fifa09.com","entry_month":"08-2025"}`)

	service := NewTenantService(mockRepo)

	tenant, err := service.CreateTenant(jsonTenant)
	if err != nil {
		t.Fatal(err)
	}

	if *tenant != *expectedTenant {
		t.Fatal("Expected tenant is different from created tenant")
	}
}

func TestTenantService_CreateTenant_Fails_DNIAlreadyExists(t *testing.T) {
	mockRepo := &mockTenantRepository{tenants: map[int]*tenant.Tenant{
		1: {
			ID:         1,
			DNI:        11111111,
			Name:       "Frodo",
			LastName:   "Baggins",
			Address:    "Unnamed road 123",
			Phone:      "+5213337778",
			Email:      "fbaggins@hobbiton.org",
			EntryMonth: calendar.NewMonthOfYear(8, 2025),
		},
	}}

	service := NewTenantService(mockRepo)

	jsonTenant := []byte(`{"dni":11111111,"name":"Bilbo","last_name":"Baggins","address":"Unnamed road 123","phone":"+5213337778","email":"bilbo@baggins.corp","entry_month":"10-2024"}`)

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

	jsonTenant := []byte(`{"dni":"hola","name":"Samwise","last_name":"Gamyi","address":"Beyond the Water 555","phone":"+5213337712","email":"sam@gamyi.com","entry_month":"06-2025"}`)

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
