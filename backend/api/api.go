package api

import (
	"cochera/config"
	"cochera/handlers"
	"fmt"
	"log"
	"net/http"
)

type API struct {
	server *http.Server
}

func New(config *config.Config) *API {
	router := setupEndpointsRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: router,
	}
	return &API{server}
}

func setupEndpointsRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc(handlers.HealthRoute, handlers.HealthHandler)
	return router
}

func (api *API) Run() {
	log.Println("🚀 API corriendo en puerto:", api.server.Addr)

	if err := api.server.ListenAndServe(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
