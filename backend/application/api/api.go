package api

import (
	tenantservice "cochera/application/services/tenant"
	"cochera/infrastructure/repositories/tenant/postgres"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

type API struct {
	server        *http.Server
	db            *pgxpool.Pool
	tenantService *tenantservice.TenantService
}

func NewAPI(port int, db *pgxpool.Pool) *API {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: nil,
	}

	return &API{server: server, db: db, tenantService: setupTenantService(db)}
}

func setupTenantService(db *pgxpool.Pool) *tenantservice.TenantService {
	repo := postgres.NewPostgresTenantRepository(db)
	return tenantservice.NewTenantService(repo)
}

func (api *API) addCORSToRouter(router http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return c.Handler(router)
}

func (api *API) Routes() *http.ServeMux {
	apiRouter := http.NewServeMux()
	apiRouter.HandleFunc(HealthRoute, api.getHealthStatus)
	apiRouter.HandleFunc(TenantsBaseRoute, api.tenantHandler)
	apiRouter.HandleFunc(TenantsBaseRoute+"/", api.tenantByIDHandler)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))

	return mainRouter
}

func (api *API) DB() *pgxpool.Pool {
	return api.db
}

func (api *API) Run() {
	defer api.db.Close()
	api.server.Handler = api.addCORSToRouter(api.Routes())

	if err := api.db.Ping(context.Background()); err != nil {
		log.Println("Failed connecting to database: ", err)
		return
	}

	port, _ := strings.CutPrefix(api.server.Addr, ":")
	log.Printf("🚀 API successfully running on port %s!", port)

	if err := api.server.ListenAndServe(); err != nil {
		log.Printf("Could not start server: %v", err)
	}
}

func (api *API) CloseDB() {
	api.db.Close()
}
