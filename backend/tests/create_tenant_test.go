package tests

import (
	"bytes"
	"cochera/api"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	api, err := api.NewTestingAPI()
	if err != nil {
		t.Fatal("Could not create testing API: ", err)
	}
	api.Run()
	defer api.Stop()

	jsonData, _ := json.Marshal(map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	})

	response, err := http.Post(api.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", api.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	jsonBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed reading response body: %v", err)
	}

	var jsonBody map[string]any
	if err := json.Unmarshal(jsonBytes, &jsonBody); err != nil {
		t.Fatalf("Failed parsing response body: %v", err)
	}

	expectedDetail := "required attributes: dni, name, last_name or entry_month"
	if receivedDetail, ok := jsonBody["detail"]; !ok || receivedDetail != expectedDetail {
		t.Fatal(expectedDetail)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code 400, got %d", response.StatusCode)
	}
}
