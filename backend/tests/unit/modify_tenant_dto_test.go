package unit

import (
	"cochera/application/dtos"
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
