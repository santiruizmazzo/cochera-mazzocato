package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"
	"net/http"
	"testing"
)

func TestFreeAlreadyTakenSlotSuccessfully_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	newTenant := domain.NewTenantBuilder().Build()
	_, _ = testAPI.CreateTenant(newTenant)

	response, err := testAPI.UpdateSlot(3, map[string]any{"tenant_id": nil})
	if err != nil {
		t.Fatalf("Failed sending PATCH request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	response, err = testAPI.GetSlotByID(3)
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

	utils.AssertResponseContains(responseMap, "id", float64(3), t)
	utils.AssertResponseContains(responseMap, "number", float64(3), t)
	utils.AssertResponseIsNil(responseMap, "tenant_id", t)

	testAPI.ClearTenants()
}
