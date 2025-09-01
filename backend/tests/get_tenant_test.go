package tests

import (
	"bytes"
	"cochera/tests/utils"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTenantThatDoesNotExist_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         123123,
		"name":        "Huang",
		"last_name":   "Lee",
		"entry_month": "07-2008",
	})

	_, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	response, err := http.Get(testApi.GetTenantsRoute() + "/999")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "tenant not found", t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)
}
