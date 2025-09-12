package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTenantsWithoutFilterSuccessfully_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

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

	response, err := testAPI.CreateTenant(firstExpectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
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

	response, err = testAPI.CreateTenant(secondExpectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	response, err = http.Get(testAPI.GetTenantsRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
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

	testAPI.ClearTenants()

	response, err := http.Get(testAPI.GetTenantsRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "no matching tenants were found", t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)
}

func TestGetTenantsFilteredByNameMatchCompletely_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	nameToFilter := "Salvatore"
	expectedTenant := domain.NewTenantBuilder().WithName(nameToFilter).Build()
	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(1).WithEmail("a@a.com").Build()
	response, err = testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(2).WithName(nameToFilter).WithEmail("b@b.com").Build()
	response, err = testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	response, err = http.Get(testAPI.GetTenantsRoute() + "?name=" + nameToFilter)
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
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

func TestGetTenantsFilteredByNameMatchPartially_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := domain.NewTenantBuilder().WithName("Martín").Build()
	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(1).WithEmail("a@a.com").Build()
	response, err = testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(2).WithName("Mario").WithEmail("b@b.com").Build()
	response, err = testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	expectedTenant = domain.NewTenantBuilder().WithDNI(3).WithName("Lamar").WithEmail("c@c.com").Build()
	response, err = testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	response, err = http.Get(testAPI.GetTenantsRoute() + "?name=" + "Mar")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	tenantsList := utils.AssertSliceOfMaps(responseMap["data"], t)

	if len(tenantsList) != 3 {
		t.Fatal("Expected a list of tenants of size 3, got ", len(tenantsList))
	}

	utils.AssertResponseStringContains(tenantsList[0]["name"], "Mar", t)
	utils.AssertResponseStringContains(tenantsList[1]["name"], "Mar", t)
	utils.AssertResponseStringContains(tenantsList[2]["name"], "Mar", t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}

func TestGetTenantsFilteredByNameDoesNotMatch_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := domain.NewTenantBuilder().Build()
	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	response, err = http.Get(testAPI.GetTenantsRoute() + "?name=" + "Agustín")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "no matching tenants were found", t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)
}

func TestGetTenantsFilteredByLastNameMatchCompletely_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	_, err = testAPI.CreateTenant(domain.NewTenantBuilder().Build())

	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	_, err = testAPI.CreateTenant(domain.NewTenantBuilder().WithDNI(1).WithEmail("a@a.com").WithLastName("Jaoming").Build())

	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	response, err := http.Get(testAPI.GetTenantsRoute() + "?lastName=" + "Jaoming")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	tenantsList := utils.AssertSliceOfMaps(responseMap["data"], t)

	if len(tenantsList) != 1 {
		t.Fatal("Expected a list of tenants of size 1, got ", len(tenantsList))
	}

	utils.AssertResponseStringContains(tenantsList[0]["last_name"], "Jaoming", t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)
}
