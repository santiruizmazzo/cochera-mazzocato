package handlers

import (
	"encoding/json"
	"net/http"
)

const HealthRoute string = "/health"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Cochera actualmente operativa!"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error al codificar respuesta en formato JSON", http.StatusInternalServerError)
	}
}
