package tests

import (
	"cochera/api"

	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

var testApi *api.TestingAPI
var err error

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()

	testApi, err = api.NewTestingAPI()
	if err != nil {
		log.Println("Could not create testing API: ", err)
		return
	}

	defer testApi.Stop()
	testApi.Run()
	code = m.Run()
}

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	})

	createTenantRoute := testApi.GetTenantCreationRoute()
	response, err := http.Post(createTenantRoute, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", createTenantRoute, err)
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

func TestCreateTenantWithDuplicateDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	setupExistingTenant(t)

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Johnny",
		"last_name":   "Klebitz",
		"entry_month": "11-2025",
	})

	createTenantRoute := testApi.GetTenantCreationRoute()
	response, err := http.Post(createTenantRoute, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", createTenantRoute, err)
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

	expectedDetail := "dni already exists"
	if receivedDetail, ok := jsonBody["detail"]; !ok || receivedDetail != expectedDetail {
		t.Fatal(expectedDetail)
	}

	if response.StatusCode != http.StatusConflict {
		t.Fatalf("Expected status code 409, got %d", response.StatusCode)
	}
}

func setupExistingTenant(t *testing.T) {
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"entry_month": "09-2024",
	})

	createTenantRoute := testApi.GetTenantCreationRoute()
	_, err := http.Post(createTenantRoute, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", createTenantRoute, err)
	}
}
