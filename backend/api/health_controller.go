package api

import (
	"cochera/internal/version"
	"encoding/json"
	"net/http"
)

const HealthRoute string = "/health"

func (api *API) getHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "operational", "version": version.Current()}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error al codificar respuesta en formato JSON", http.StatusInternalServerError)
	}
}
