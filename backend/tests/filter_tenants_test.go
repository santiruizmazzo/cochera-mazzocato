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
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()

	firstExpectedTenant := map[string]any{
		"id":          float64(1),
		"dni":         float64(123),
		"name":        "Hsin",
		"last_name":   "Jaoming",
		"entry_month": "07-2008",
		"address":     nil,
		"email":       nil,
		"phone":       nil,
	}
	jsonData, _ := json.Marshal(firstExpectedTenant)

	_, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	secondExpectedTenant := map[string]any{
		"id":          float64(2),
		"dni":         float64(321),
		"name":        "Chan",
		"last_name":   "Jaoming",
		"entry_month": "07-2008",
		"address":     nil,
		"email":       nil,
		"phone":       nil,
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
	tenantsList := utils.AssertSliceOfMaps(t, responseMap["data"])

	if !reflect.DeepEqual(tenantsList[0], firstExpectedTenant) {
		t.Fatalf("Expected %+v, got %+v", firstExpectedTenant, tenantsList[0])
	}
	if !reflect.DeepEqual(tenantsList[1], secondExpectedTenant) {
		t.Fatalf("Expected %+v, got %+v", secondExpectedTenant, tenantsList[1])
	}

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
