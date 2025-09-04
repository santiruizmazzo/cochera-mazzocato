package main

import (
	"cochera/application/api"
	"cochera/internal/config"
	"log"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Failed loading configuration: ", err)
	}

	api := api.NewAPI(config.Port, config.DB)
	api.Run()
}
