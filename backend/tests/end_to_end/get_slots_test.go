package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"
	"net/http"
	"testing"
)

func TestGetAllSlots_EndToEnd(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	tenant := domain.NewTenantBuilder().Build()
	_, err = testAPI.CreateTenant(tenant)
	if err != nil {
		t.Fatal(err)
	}

	tenant = domain.NewTenantBuilder().Build()
	_, err = testAPI.CreateTenant(tenant)
	if err != nil {
		t.Fatal(err)
	}

	response, err := testAPI.UpdateSlot(1, map[string]int{"tenant_id": 1})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.UpdateSlot(2, map[string]int{"tenant_id": 1})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.UpdateSlot(3, map[string]int{"tenant_id": 1})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.UpdateSlot(4, map[string]int{"tenant_id": 1})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.UpdateSlot(5, map[string]int{"tenant_id": 2})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.UpdateSlot(6, map[string]int{"tenant_id": 2})
	if err != nil {
		t.Fatal(err)
	}

	response, err = testAPI.GetSlots()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	slotsList := utils.AssertSliceOfMaps(responseMap["data"], t)

	if len(slotsList) != 12 {
		t.Fatal("Expected a list of slots of size 12, got ", len(slotsList))
	}

	utils.AssertResponseContains(slotsList[0], "tenant_id", float64(1), t)
	utils.AssertResponseContains(slotsList[1], "tenant_id", float64(1), t)
	utils.AssertResponseContains(slotsList[2], "tenant_id", float64(1), t)
	utils.AssertResponseContains(slotsList[3], "tenant_id", float64(1), t)
	utils.AssertResponseContains(slotsList[4], "tenant_id", float64(2), t)
	utils.AssertResponseContains(slotsList[5], "tenant_id", float64(2), t)

	utils.AssertResponseIsNil(slotsList[6], "tenant_id", t)
	utils.AssertResponseIsNil(slotsList[7], "tenant_id", t)
	utils.AssertResponseIsNil(slotsList[8], "tenant_id", t)
	utils.AssertResponseIsNil(slotsList[9], "tenant_id", t)
	utils.AssertResponseIsNil(slotsList[10], "tenant_id", t)
	utils.AssertResponseIsNil(slotsList[11], "tenant_id", t)

	testAPI.ClearTenants()
}
