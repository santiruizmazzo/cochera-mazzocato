package tests

import (
	"bytes"
	"cochera/tests/utils"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTenantsWithoutFilterSuccessfully_EndToEnd(t *testing.T) {
	t.Skip("Skipping this test temporarily...")

	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()

	firstExpectedTenant := map[string]any{
		"dni":         123,
		"name":        "Hsin",
		"last_name":   "Jaoming",
		"entry_month": "07-2008",
	}
	jsonData, _ := json.Marshal(firstExpectedTenant)

	_, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	secondExpectedTenant := map[string]any{
		"dni":         321,
		"name":        "Chan",
		"last_name":   "Jaoming",
		"entry_month": "07-2008",
	}
	jsonData, _ = json.Marshal(secondExpectedTenant)

	_, err = http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	response, err := http.Get(testApi.GetTenantsRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	switch value := responseMap["data"].(type) {
	case []map[string]any:
		if !reflect.DeepEqual(value[0], firstExpectedTenant) {
			t.Fatalf("Expected %s, got %s", firstExpectedTenant, value[0])
		}
		if !reflect.DeepEqual(value[1], secondExpectedTenant) {
			t.Fatalf("Expected %s, got %s", secondExpectedTenant, value[1])
		}
	default:
		t.Fatal("ERRORRRRRRR")
	}

	// utils.AssertResponseContains(responseMap, "id", float64(tenantID), t)
	// utils.AssertResponseContains(responseMap, "name", "Wu 'Kenny'", t)
	// utils.AssertResponseContains(responseMap, "last_name", "Lee", t)
	// utils.AssertResponseContains(responseMap, "entry_month", "07-2008", t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
