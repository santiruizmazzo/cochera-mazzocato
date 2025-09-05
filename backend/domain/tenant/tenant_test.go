package tenant

import (
	myerrors "cochera/domain/errors"
	"encoding/json"
	"errors"
	"testing"
)

func TestTenantCreatedFromJSON(t *testing.T) {
	expectedTenant := NewTenantBuilder().Build()
	jsonTenant, _ := json.Marshal(expectedTenant)

	tenant, err := NewTenantFromJSON(jsonTenant)
	if err != nil {
		t.Fatal("Failed creating tenant from json: ", err)
	}

	if *expectedTenant != *tenant {
		t.Fatalf("Expected %v, got %v", expectedTenant, tenant)
	}
}

func TestNewTenantFromJSONReturnsCustomErrorWhenGivenNonNumericDNI(t *testing.T) {
	jsonTenant, _ := json.Marshal(map[string]any{
		"dni":         "adios",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	})

	tenant, err := NewTenantFromJSON(jsonTenant)

	if tenant != nil || err == nil {
		t.Fatal("Tenant creation from json should fail when DNI is not a number")
	}

	if !errors.Is(err, myerrors.ErrDNIMustBeNumber) {
		t.Fatal("Returned error should be of type ErrInvalidDNI")
	}
}
