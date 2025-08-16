package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	server *http.Server
	db     *pgxpool.Pool
}

func NewAPI(port int, db *pgxpool.Pool) *API {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: nil,
	}

	return &API{server: server, db: db}
}

func (api *API) Routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc(HealthRoute, api.healthHandler)
	return router
}

func (api *API) Run() {
	defer api.db.Close()
	api.server.Handler = api.Routes()

	if err := api.db.Ping(context.Background()); err != nil {
		log.Println("Error al conectar con la base de datos: ", err)
		return
	}
	log.Println("🚀 API corriendo en puerto:", api.server.Addr)

	if err := api.server.ListenAndServe(); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}

func (api *API) CloseDB() {
	api.db.Close()
}
