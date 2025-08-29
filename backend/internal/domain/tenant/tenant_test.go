package tenant

import (
	"cochera/internal/domain/time"
	"fmt"
	"testing"
)

func TestTenantCreatedFromJSON(t *testing.T) {
	const expectedDni uint32 = 43295798
	const expectedName string = "Santiago"
	const expectedLastName string = "Ruiz Mazzocato"
	const expectedAddress string = "Roseti 745"
	const expectedPhone string = "543442407277"
	const expectedEmail string = "santimazzo98@gmail.com"
	const expectedEntryMonth string = "05-2025"

	jsonString := fmt.Sprintf(`{"dni":%v,"name":"%s","last_name":"%s","address":"%s","phone":"%s","email":"%s","entry_month":"%s"}`, expectedDni, expectedName, expectedLastName, expectedAddress, expectedPhone, expectedEmail, expectedEntryMonth)
	tenant, err := NewTenantFromJSON([]byte(jsonString))
	if err != nil {
		t.Fatal("Failed creating tenant from json: ", err)
	}

	if tenant.DNI != expectedDni {
		t.Fatalf("Expected DNI %v, got %v", expectedDni, tenant.DNI)
	}

	if tenant.Name != expectedName {
		t.Fatalf("Expected name %s, got %s", expectedName, tenant.Name)
	}

	if tenant.LastName != expectedLastName {
		t.Fatalf("Expected last name %s, got %s", expectedLastName, tenant.LastName)
	}

	if tenant.Address != expectedAddress {
		t.Fatalf("Expected address %s, got %s", expectedAddress, tenant.Address)
	}

	if tenant.Phone != expectedPhone {
		t.Fatalf("Expected phone %s, got %s", expectedPhone, tenant.Phone)
	}

	if tenant.Email != expectedEmail {
		t.Fatalf("Expected email %s, got %s", expectedEmail, tenant.Email)
	}

	expectedMonthOfYear, err := time.NewMonthOfYearFromString(expectedEntryMonth)
	if tenant.EntryMonth != expectedMonthOfYear || err != nil {
		t.Fatalf("Expected entry month %s, got %s", expectedEntryMonth, tenant.EntryMonth)
	}
}
