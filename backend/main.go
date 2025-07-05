package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	puerto := os.Getenv("PUERTO")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "La cochera esta operativa actualmente!"}
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", puerto), nil)
}
