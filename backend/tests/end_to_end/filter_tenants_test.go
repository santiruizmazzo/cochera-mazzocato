package endtoend

import (
	"bytes"
	"cochera/domain"
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
	tenantsList := utils.AssertSliceOfMaps(responseMap["data"], t)

	if !reflect.DeepEqual(tenantsList[0], firstExpectedTenant) {
		t.Fatalf("Expected %+v, got %+v", firstExpectedTenant, tenantsList[0])
	}
	if !reflect.DeepEqual(tenantsList[1], secondExpectedTenant) {
		t.Fatalf("Expected %+v, got %+v", secondExpectedTenant, tenantsList[1])
	}

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}

func TestGetTenantsWithoutAnyTenantsCreated_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()

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

	utils.AssertResponseContains(responseMap, "detail", "there are no tenants created", t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)
}

func TestGetTenantsFilteredByNameMatchCompletely_EndToEnd(t *testing.T) {
	t.Skip("Skipping this test temporarily...")

	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()

	nameToFilter := "Salvatore"
	expectedTenant := domain.NewTenantBuilder().WithName(nameToFilter).Build()
	jsonTenant, _ := json.Marshal(expectedTenant)
	_, err = http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(1).WithEmail("a@a.com").Build()
	jsonTenant, _ = json.Marshal(expectedTenant)
	_, err = http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(2).WithName(nameToFilter).WithEmail("b@b.com").Build()
	jsonTenant, _ = json.Marshal(expectedTenant)
	_, err = http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	response, err := http.Get(testApi.GetTenantsRoute() + "?name=" + nameToFilter)
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	tenantsList := utils.AssertSliceOfMaps(responseMap["data"], t)

	if len(tenantsList) != 2 {
		t.Fatal("Expected a list of tenants of size 2, got ", len(tenantsList))
	}

	utils.AssertResponseContains(tenantsList[0], "name", nameToFilter, t)
	utils.AssertResponseContains(tenantsList[1], "name", nameToFilter, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
