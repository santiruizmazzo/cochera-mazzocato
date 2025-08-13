package tests

import (
	"cochera/api"
	"io"
	"net/http"
	"strings"
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error leyendo response body: %v", err)
	}

	responseJSON := strings.TrimSpace(string(bodyBytes))
	expectedJSON := `{"message":"Cochera actualmente operativa!"}`
	if responseJSON != expectedJSON {
		t.Errorf("JSON esperado: '%s', obtenido '%s'", expectedJSON, responseJSON)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado codigo 200, obtenido %d", resp.StatusCode)
	}
}
