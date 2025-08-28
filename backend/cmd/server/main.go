package main

import (
	"cochera/api"
	"cochera/internal/config"
	"log"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error cargando configuración: ", err)
	}

	api := api.NewAPI(config.Port, config.DB)
	api.Run()
}
