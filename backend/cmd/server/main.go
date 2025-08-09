package main

import (
	"cochera/config"
	"cochera/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error al cargar configuracion desde archivo .env")
	}

	http.HandleFunc("/health", handlers.HealthHandler)

	fmt.Printf("🚀 Servidor ejecutándose en puerto %v\n", config.Port)

	address := fmt.Sprintf(":%v", config.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
