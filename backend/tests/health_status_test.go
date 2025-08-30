package tests

import (
	"cochera/internal/version"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestHealthStatusEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	response, err := http.Get(testApi.GetHealthStatusRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetHealthStatusRoute(), err)
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

	if expectedStatus, ok := jsonBody["status"]; !ok || expectedStatus != "operational" {
		t.Fatalf("Status not found, or does not match with expected")
	}

	if expectedVersion, ok := jsonBody["version"]; !ok || expectedVersion != version.Current() {
		t.Fatalf("Version not found, or does not match with expected")
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", response.StatusCode)
	}
}
