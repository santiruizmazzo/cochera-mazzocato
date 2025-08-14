package main

import (
	"cochera/api"
	"cochera/config"
	"log"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error cargando configuración: ", err)
	}

	api, err := api.New(config)
	if err != nil {
		log.Fatal("Error conectando a base de datos: ", err)
	}
	api.Run()
}
