package main

import (
	"cochera/api"
	"cochera/config"
	"log"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error cargando configuración del servidor")
	}

	api := api.New(config)
	api.Run()
}
