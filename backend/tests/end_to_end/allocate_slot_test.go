package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"
	"net/http"
	"testing"
)

func TestAllocateFreeSlotToTenant_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	newTenant := domain.NewTenantBuilder().Build()
	_, _ = testAPI.CreateTenant(newTenant)

	requestBody := map[string]int{
		"tenant_id": 1,
	}

	response, err := testAPI.UpdateSlot(1, requestBody)
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	response, err = http.Get(testAPI.GetSlotsRoute() + "/1")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatal("Failed closing response body: ", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "id", float64(1), t)
	utils.AssertResponseContains(responseMap, "number", float64(1), t)
	utils.AssertResponseContains(responseMap, "tenant_id", float64(1), t)

	testAPI.ClearTenants()
}

func TestAllocateNonExistingSlotToTenant_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	tenant := domain.NewTenantBuilder().Build()
	_, err = testAPI.CreateTenant(tenant)
	if err != nil {
		t.Fatal(err)
	}

	response, err := testAPI.UpdateSlot(13, map[string]int{"tenant_id": 1})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusNotFound, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "plaza no encontrada", t)

	testAPI.ClearTenants()
}

func TestAllocateSlotToNonExistingTenant_EndToEnd(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	response, err := testAPI.UpdateSlot(5, map[string]int{"tenant_id": 5})
	if err != nil {
		t.Fatal(err)
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
