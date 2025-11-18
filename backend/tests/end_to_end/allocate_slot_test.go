package endtoend

import (
	"bytes"
	"cochera/tests/utils"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAllocateFreeSlotToTenant_EndToEnd(t *testing.T) {
	t.Skip("SKIPEANDOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")

	if testing.Short() {
		t.Skip()
	}

	jsonBody, _ := json.Marshal(map[string]int{
		"tenant_id": 1,
		"slot_id":   1,
	})

	response, err := utils.HTTPPut(testAPI.GetSlotsRoute(), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed sending PUT request to %s: %v", testAPI.GetSlotsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "slot_id", 1, t)

	utils.AssertResponseContains(responseMap, "tenant_id", 1, t)
}
