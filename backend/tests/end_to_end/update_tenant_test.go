package endtoend

import (
	"cochera/tests/utils"

	"net/http"
	"testing"
)

func TestUpdateTenantSuccessfully_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	expectedTenant := map[string]any{
		"dni":         14571272,
		"name":        "Victor",
		"last_name":   "Vance",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant["dni"] = 43295798
	expectedTenant["name"] = "Lance"
	expectedTenant["entry_month"] = "02-2024"
	modifiedTenant := expectedTenant

	response, err = testAPI.UpdateTenant(1, modifiedTenant)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "dni", float64(43295798), t)
	utils.AssertResponseContains(responseMap, "name", "Lance", t)
	utils.AssertResponseContains(responseMap, "entry_month", "02-2024", t)
}

func TestUpdateTenantThatDoesNotExist_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	modifiedTenant := map[string]any{
		"dni":         99999999,
		"name":        "Lance",
		"last_name":   "Vance",
		"entry_month": "01-2025",
	}

	response, err := testAPI.UpdateTenant(66666, modifiedTenant)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "inquilino no encontrado", t)
}

func TestUpdateTenantFailsWhenChangingToInvalidEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := map[string]any{
		"dni":         14571272,
		"name":        "Juan Alberto",
		"last_name":   "García",
		"entry_month": "01-2025",
		"email":       "juanalberto@garcia.com",
	}

	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatal("Failed creating tenant: ", err)
	}

	modifiedTenant := map[string]any{
		"email": "illojuan.com",
	}

	response, err = testAPI.UpdateTenant(1, modifiedTenant)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el email debe seguir el formato estándar", t)
}

func TestUpdateTenantFailsWhenDeletingRequiredAttribute_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := map[string]any{
		"dni":         14571272,
		"name":        "Juan Alberto",
		"last_name":   "García",
		"entry_month": "01-2025",
		"email":       "juanalberto@garcia.com",
	}

	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatal("Failed creating tenant: ", err)
	}

	modifiedTenant := map[string]any{
		"dni": nil,
	}

	response, err = testAPI.UpdateTenant(1, modifiedTenant)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI debe ser un número entero", t)
}

func TestUpdateTenantDoesNotNullifyValueOfMissingAttributeInRequestBody_EndToEnd(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := map[string]any{
		"dni":         14571272,
		"name":        "Juan Alberto",
		"last_name":   "García",
		"entry_month": "01-2025",
		"email":       "juanalberto@garcia.com",
		"address":     "Calle Larios 123",
		"phone":       "+5494444192929",
	}

	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatal("Failed creating tenant: ", err)
	}

	modifiedTenant := map[string]any{
		"dni": 55,
	}

	response, err = testAPI.UpdateTenant(1, modifiedTenant)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "dni", float64(55), t)
	utils.AssertResponseContains(responseMap, "name", "Juan Alberto", t)
	utils.AssertResponseContains(responseMap, "last_name", "García", t)
	utils.AssertResponseContains(responseMap, "entry_month", "01-2025", t)
	utils.AssertResponseContains(responseMap, "email", "juanalberto@garcia.com", t)
	utils.AssertResponseContains(responseMap, "address", "Calle Larios 123", t)
	utils.AssertResponseContains(responseMap, "phone", "+5494444192929", t)
}
