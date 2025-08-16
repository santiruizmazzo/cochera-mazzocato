package tests

import (
	"cochera/api"
	"cochera/version"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestHealthStatus(t *testing.T) {
	api := api.NewMock()
	api.Run()
	defer api.Stop()

	resp, err := http.Get(api.GetHealthFullRoute())
	if err != nil {
		t.Errorf("Error al llamar a %s: %v", api.GetHealthFullRoute(), err)
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			t.Errorf("Error cerrando el response body: %v", cerr)
		}
	}()

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error leyendo response body: %v", err)
	}

	var jsonBody map[string]any
	if err := json.Unmarshal(jsonBytes, &jsonBody); err != nil {
		t.Errorf("Error parseando response body: %v", err)
	}

	if expectedStatus, ok := jsonBody["status"]; !ok || expectedStatus != "operational" {
		t.Errorf("Status no encontrado, o no coincide con esperado")
	}

	if expectedVersion, ok := jsonBody["version"]; !ok || expectedVersion != version.Current() {
		t.Errorf("Version no encontrada, o no coincide con esperada")
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado código 200, obtenido %d", resp.StatusCode)
	}
}
