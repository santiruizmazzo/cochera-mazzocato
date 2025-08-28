package tests

import (
	"bytes"
	"cochera/api"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	api, err := api.NewTestingAPI()
	if err != nil {
		t.Error("Error al instanciar servidor para testing: ", err)
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
		t.Fatalf("Error al llamar a %s: %v", api.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Error cerrando el response body: %v", cerr)
		}
	}()

	jsonBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error leyendo response body: %v", err)
	}

	var jsonBody map[string]any
	log.Println(string(jsonBytes))
	if err := json.Unmarshal(jsonBytes, &jsonBody); err != nil {
		t.Fatalf("Error parseando response body: %v", err)
	}

	expectedDetail := "required attributes: dni, name, last_name or entry_month"
	if receivedDetail, ok := jsonBody["detail"]; !ok || receivedDetail != expectedDetail {
		t.Fatal(expectedDetail)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("Esperado código 400, obtenido %d", response.StatusCode)
	}
}
