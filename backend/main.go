package main

import (
	"cochera/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Cochera actualmente operativa!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error al cargar configuracion desde archivo .env")
	}

	http.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), nil))
}
