package endtoend

import (
	"cochera/tests/utils"

	"net/http"
	"testing"
)

func TestModifyTenantSuccessfully_EndToEnd(t *testing.T) {
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
	modifiedTenant := expectedTenant

	response, err = testAPI.ModifyTenant(1, modifiedTenant)
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
}

func TestModifyTenantThatDoesNotExist_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	modifiedTenant := map[string]any{
		"dni":         99999999,
		"name":        "Lance",
		"last_name":   "Vance",
		"entry_month": "01-2025",
	}

	response, err := testAPI.ModifyTenant(66666, modifiedTenant)
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
