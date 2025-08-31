package tests

import (
	"cochera/tests/utils"

	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func setupExistingTenant(t *testing.T) {
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"entry_month": "09-2024",
	})

	_, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}
}

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "dni is required", t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)
}

func TestCreateTenantWithDuplicateDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	setupExistingTenant(t)

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Johnny",
		"last_name":   "Klebitz",
		"entry_month": "11-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusConflict, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni already exists", t)
}

func TestCreateTenantWithNonNumericDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         "17888420",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni must be a positive number", t)
}
