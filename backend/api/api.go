package api

import (
	"cochera/config"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	server *http.Server
	DB     *pgxpool.Pool
}

func New(config *config.Config) (*API, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: nil,
	}
	db, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &API{server: server, DB: db}, nil
}

func (api *API) setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc(HealthRoute, api.healthHandler)
	return router
}

func (api *API) Run() {
	defer api.DB.Close()
	api.server.Handler = api.setupRouter()
	log.Println("🚀 API corriendo en puerto:", api.server.Addr)

	if err := api.server.ListenAndServe(); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
