package endtoend

import (
	"bytes"
	"cochera/domain"
	"cochera/tests/utils"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestAllocateFreeSlotToTenant_EndToEnd(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	newTenant := domain.NewTenantBuilder().Build()
	_, _ = testAPI.CreateTenant(newTenant)

	jsonBody, _ := json.Marshal(map[string]int{
		"tenant_id": 1,
	})

	response, err := utils.HTTPPatch(testAPI.GetSlotsRoute()+"/1", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	response, err = http.Get(testAPI.GetSlotsRoute() + "/1")
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)
	log.Println(responseMap)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "id", 1, t)
	utils.AssertResponseContains(responseMap, "number", 1, t)
	utils.AssertResponseContains(responseMap, "tenant_id", 1, t)

	testAPI.ClearTenants()
}
