package endtoend

import (
	"bytes"
	"cochera/tests/utils"
	"encoding/json"

	"net/http"
	"testing"
)

func TestModifyTenantSuccessfully_EndToEnd(t *testing.T) {
	t.Skip()

	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

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

	jsonBody, _ := json.Marshal(map[string]any{
		"dni": 43295798,
	})

	fullEndpointURL := testAPI.GetTenantsRoute() + "/1"

	response, err = utils.HTTPPatch(fullEndpointURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
