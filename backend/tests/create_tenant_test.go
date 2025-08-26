package tests

import (
	"cochera/api"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCreateTenantWithMissingAttributesEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	api, err := api.NewTestingAPI()
	if err != nil {
		t.Error("Error al instanciar servidor para testing: ", err)
	}
	api.Run()
	defer api.Stop()

	response, err := http.Get(api.GetTenantCreationRoute())
	if err != nil {
		t.Errorf("Error al llamar a %s: %v", api.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Errorf("Error cerrando el response body: %v", cerr)
		}
	}()

	jsonBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error leyendo response body: %v", err)
	}

	var jsonBody map[string]any
	if err := json.Unmarshal(jsonBytes, &jsonBody); err != nil {
		t.Errorf("Error parseando response body: %v", err)
	}

	expectedDetail := "falta alguno de estos datos: DNI, nombre o apellido"
	if receivedDetail, ok := jsonBody["detail"]; !ok || receivedDetail != expectedDetail {
		t.Error(expectedDetail)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Esperado código 400, obtenido %d", response.StatusCode)
	}
}
