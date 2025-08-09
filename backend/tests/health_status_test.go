package tests

import (
	"cochera/handlers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthStatus(t *testing.T) {
	const healthEndpoint = "/health"

	mux := http.NewServeMux()
	mux.HandleFunc(healthEndpoint, handlers.HealthHandler)
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + healthEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			t.Errorf("Error cerrando el response body: %v", cerr)
		}
	}()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error leyendo el response body: %v", err)
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
