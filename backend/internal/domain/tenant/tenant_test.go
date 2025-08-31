package tenant

import (
	"cochera/internal/domain/calendar"
	myerrors "cochera/internal/errors"
	"encoding/json"
	"errors"
	"testing"
)

func TestTenantCreatedFromJSON(t *testing.T) {
	expectedTenant := &Tenant{
		DNI:        43295798,
		Name:       "Santiago",
		LastName:   "Ruiz Mazzocato",
		Address:    "Roseti 745",
		Phone:      "+543442407277",
		Email:      "santimazzo98@gmail.com",
		EntryMonth: calendar.NewMonthOfYear(05, 2025),
	}
	jsonData, _ := json.Marshal(expectedTenant)

	tenant, err := NewTenantFromJSON(jsonData)
	if err != nil {
		t.Fatal("Failed creating tenant from json: ", err)
	}

	if tenant.DNI != expectedTenant.DNI {
		t.Fatalf("Expected DNI %v, got %v", expectedTenant.DNI, tenant.DNI)
	}

	if tenant.Name != expectedTenant.Name {
		t.Fatalf("Expected name %s, got %s", expectedTenant.Name, tenant.Name)
	}

	if tenant.LastName != expectedTenant.LastName {
		t.Fatalf("Expected last name %s, got %s", expectedTenant.LastName, tenant.LastName)
	}

	if tenant.Address != expectedTenant.Address {
		t.Fatalf("Expected address %s, got %s", expectedTenant.Address, tenant.Address)
	}

	if tenant.Phone != expectedTenant.Phone {
		t.Fatalf("Expected phone %s, got %s", expectedTenant.Phone, tenant.Phone)
	}

	if tenant.Email != expectedTenant.Email {
		t.Fatalf("Expected email %s, got %s", expectedTenant.Email, tenant.Email)
	}

	if tenant.EntryMonth != expectedTenant.EntryMonth || err != nil {
		t.Fatalf("Expected entry month %s, got %s", expectedTenant.EntryMonth, tenant.EntryMonth)
	}
}

func TestNewTenantFromJSONReturnsCustomErrorWhenGivenNonNumericDNI(t *testing.T) {
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         "adios",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	})

	tenant, err := NewTenantFromJSON(jsonData)

	if tenant != nil || err == nil {
		t.Fatal("Tenant creation from json should fail when DNI is not a number")
	}

	if !errors.Is(err, myerrors.ErrDNIMustBeNumber) {
		t.Fatal("Returned error should be of type ErrInvalidDNI")
	}
}
