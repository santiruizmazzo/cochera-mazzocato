package endtoend

import (
	"bytes"
	"cochera/tests/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetTenantByIDThatDoesNotExist_EndToEnd(t *testing.T) {
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

func TestGetTenantByIDSuccessfully_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         171616,
		"name":        "Wu 'Kenny'",
		"last_name":   "Lee",
		"entry_month": "07-2008",
	})

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	responseMap := utils.CreateMapFromBody(response.Body, t)
	var getTenantByIDRoute string
	var tenantID int

	switch value := responseMap["id"].(type) {
	case float64:
		tenantID = int(value)
		getTenantByIDRoute = testApi.GetTenantsRoute() + "/" + fmt.Sprint(tenantID)
	default:
		t.Fatal("Error reading response to creation of tenant")
	}

	response, err = http.Get(getTenantByIDRoute)
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap = utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "id", float64(tenantID), t)
	utils.AssertResponseContains(responseMap, "name", "Wu 'Kenny'", t)
	utils.AssertResponseContains(responseMap, "last_name", "Lee", t)
	utils.AssertResponseContains(responseMap, "entry_month", "07-2008", t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
