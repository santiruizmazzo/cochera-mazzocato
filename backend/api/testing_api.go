package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TestingAPI struct {
	server  *httptest.Server
	router  *http.ServeMux
	realAPI *API
}

func NewTestingAPI() (*TestingAPI, error) {
	databaseURL := os.Getenv("TEST_DB_URL")
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	realAPI := NewAPI(0, db)
	return &TestingAPI{router: realAPI.Routes(), realAPI: realAPI}, nil
}

func (api *TestingAPI) Run() {
	api.server = httptest.NewServer(api.router)
}

func (api *TestingAPI) Stop() {
	api.realAPI.CloseDB()
	api.server.Close()
}

func (api *TestingAPI) GetHealthFullRoute() string {
	return api.server.URL + "/health"
}

func (api *TestingAPI) GetTenantCreationRoute() string {
	return api.server.URL + "/tenants"
}
