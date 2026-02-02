package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAllSlots_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	tenant := domain.NewTenantBuilder().Build()
	_, err = testAPI.CreateTenant(tenant)
	if err != nil {
		t.Fatal(err)
	}

	tenant = domain.NewTenantBuilder().WithDNI(7).WithEmail("caca@caca.com").Build()
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

	expectedResponseMap := map[string]any{
		"data": []map[string]any{
			{
				"id":     float64(1),
				"number": float64(1),
				"tenant": map[string]any{
					"id":        float64(1),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(2),
				"number": float64(2),
				"tenant": map[string]any{
					"id":        float64(1),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(3),
				"number": float64(3),
				"tenant": map[string]any{
					"id":        float64(1),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(4),
				"number": float64(4),
				"tenant": map[string]any{
					"id":        float64(1),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(5),
				"number": float64(5),
				"tenant": map[string]any{
					"id":        float64(2),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(6),
				"number": float64(6),
				"tenant": map[string]any{
					"id":        float64(2),
					"name":      "Huang",
					"last_name": "Lee",
				},
			},
			{
				"id":     float64(7),
				"number": float64(7),
				"tenant": nil,
			},
			{
				"id":     float64(8),
				"number": float64(8),
				"tenant": nil,
			},
			{
				"id":     float64(9),
				"number": float64(9),
				"tenant": nil,
			},
			{
				"id":     float64(10),
				"number": float64(10),
				"tenant": nil,
			},
			{
				"id":     float64(11),
				"number": float64(11),
				"tenant": nil,
			},
			{
				"id":     float64(12),
				"number": float64(12),
				"tenant": nil,
			},
		},
	}

	receivedJson, _ := json.Marshal(responseMap)
	expectedJson, _ := json.Marshal(expectedResponseMap)

	if !reflect.DeepEqual(receivedJson, expectedJson) {
		t.Fatalf("Expected %+v, got %+v", receivedJson, expectedJson)
	}

	testAPI.ClearTenants()
}
