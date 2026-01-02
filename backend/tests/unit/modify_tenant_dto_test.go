package unit

import (
	"cochera/application/dtos"
	vo "cochera/domain/value_objects"
	"encoding/json"
	"testing"
)

func TestUnmarshallingEmptyAttributes(t *testing.T) {
	var tenantData dtos.ModifyTenantDTO
	requestBody := `{"dni":12424}`

	err := json.Unmarshal([]byte(requestBody), &tenantData)

	if err != nil {
		t.Fatal("unmarshalling of input tenant dto should not fail: ", err)
	}

	if *tenantData.DNI != 12424 {
		t.Fatal("expected DNI value of 12424, got", tenantData.DNI)
	}

	if tenantData.Name != nil {
		t.Fatal("name is expected to be nil")
	}

	if tenantData.LastName != nil {
		t.Fatal("last name is expected to be nil")
	}

	if tenantData.EntryMonth != nil {
		t.Fatal("entry month is expected to be nil")
	}

	if tenantData.Address != nil {
		t.Fatal("address is expected to be nil")
	}

	if tenantData.Phone != nil {
		t.Fatal("phone is expected to be nil")
	}

	if tenantData.Email != nil {
		t.Fatal("email is expected to be nil")
	}
}

func TestUnmarshallingComplexAttributeTypesSuccessfully(t *testing.T) {
	var tenantData dtos.ModifyTenantDTO
	requestBody := `{"dni":90, "entry_month":"01-2026"}`

	err := json.Unmarshal([]byte(requestBody), &tenantData)

	if err != nil {
		t.Fatal("unmarshalling of input tenant dto should not fail: ", err)
	}

	if *tenantData.DNI != 90 {
		t.Fatal("expected DNI value of 90, got", tenantData.DNI)
	}

	expectedEntryMonth := vo.MonthOfYear{Month: 1, Year: 2026}
	if *tenantData.EntryMonth != expectedEntryMonth {
		t.Fatal("expected entry month 01-2026, got", *tenantData.EntryMonth)
	}

	if tenantData.Name != nil {
		t.Fatal("name is expected to be nil")
	}

	if tenantData.LastName != nil {
		t.Fatal("last name is expected to be nil")
	}

	if tenantData.Address != nil {
		t.Fatal("address is expected to be nil")
	}

	if tenantData.Phone != nil {
		t.Fatal("phone is expected to be nil")
	}

	if tenantData.Email != nil {
		t.Fatal("email is expected to be nil")
	}
}
