package api

import (
	"cochera/internal/version"
	"encoding/json"
	"net/http"
)

const HealthRoute string = "/health"

func (api *API) getHealthStatus(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "operational", "version": version.Current()}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed trying to encode response as json", http.StatusInternalServerError)
	}
}
