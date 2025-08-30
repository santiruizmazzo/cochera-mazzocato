package testingapi

import (
	"cochera/api"
	"cochera/tests/utils"
	"context"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TestingAPI struct {
	server  *httptest.Server
	router  *http.ServeMux
	realAPI *api.API
}

func NewTestingAPI() (*TestingAPI, error) {
	databaseURL := os.Getenv("TEST_DB_URL")
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	realAPI := api.NewAPI(0, db)
	return &TestingAPI{router: realAPI.Routes(), realAPI: realAPI}, nil
}

func (api *TestingAPI) Run() {
	api.server = httptest.NewServer(api.router)
}

func (api *TestingAPI) Stop() {
	utils.CleanupAndCloseTestDatabase(api.realAPI.DB())
	api.server.Close()
}

func (api *TestingAPI) GetHealthStatusRoute() string {
	return api.server.URL + "/health"
}

func (api *TestingAPI) GetTenantCreationRoute() string {
	return api.server.URL + "/tenants"
}
