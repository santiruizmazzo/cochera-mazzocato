package unit

import (
	"cochera/application/dtos"
	vo "cochera/domain/value_objects"
	"encoding/json"
	"testing"
)

func TestUnmarshallingEmptyAttributes(t *testing.T) {

	mapTenant := map[string]any{
		"dni": 12424,
	}
	requestBody, _ := json.Marshal(mapTenant)

	var tenantData dtos.InputTenantDTO

	err := json.Unmarshal(requestBody, &tenantData)

	if err != nil {
		t.Fatal("unmarshalling of input tenant dto should not fail: ", err)
	}

	if tenantData.DNI != 12424 {
		t.Fatal("expected DNI value of 12424, got", tenantData.DNI)
	}

	if tenantData.Name != "" {
		t.Fatal("name is expected to be nil")
	}

	if tenantData.LastName != "" {
		t.Fatal("last name is expected to be nil")
	}

	emptyMonthOfYear := vo.MonthOfYear{Month: 0, Year: 0}
	if tenantData.EntryMonth != emptyMonthOfYear {
		t.Fatal("entry month is expected to be nil")
	}

	if tenantData.Address != "" {
		t.Fatal("address is expected to be nil")
	}

	emptyPhone := vo.Phone{CountryCode: "", LineNumber: ""}
	if tenantData.Phone != emptyPhone {
		t.Fatal("phone is expected to be nil")
	}

	if tenantData.Email != "" {
		t.Fatal("email is expected to be nil")
	}
}
