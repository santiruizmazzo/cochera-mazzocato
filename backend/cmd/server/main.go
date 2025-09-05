package main

import (
	"cochera/application"
	"cochera/application/api"
	"log"
)

func main() {
	config, err := application.NewConfig()
	if err != nil {
		log.Fatal("Failed loading configuration: ", err)
	}

	api := api.NewAPI(config.Port, config.DB)
	api.Run()
}
