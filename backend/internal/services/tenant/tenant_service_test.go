package tenantservice

import (
	"cochera/internal/domain/tenant"
	"cochera/internal/domain/time"
	"testing"
)

type mockTenantRepository struct {
	tenants map[int]*tenant.Tenant
	err     error
}

func (mockRepo *mockTenantRepository) Save(tenant *tenant.Tenant) (*tenant.Tenant, error) {
	if mockRepo.err != nil {
		return nil, mockRepo.err
	}
	tenant.ID = 1
	return tenant, nil
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
		EntryMonth: time.NewMonthOfYear(8, 2025),
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
