package testingapi

import (
	"bytes"
	"cochera/application/api"
	"cochera/tests/utils"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TestingAPI is a testing-focused wrapper for the API type.
//
// Used in end to end tests to simulate a production-like environment for testing available endpoints.
type TestingAPI struct {
	server  *httptest.Server
	router  *http.ServeMux
	realAPI *api.API
}

// NewTestingAPI creates a TestingAPI instance which connects to a real database under the hood.
//
// It connects to a testing database descripted by the TEST_DB_URL environment variable.
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

func (api *TestingAPI) ClearTenants() {
	utils.ClearTenantsTable(api.realAPI.DB())
}

func (api *TestingAPI) GetHealthStatusRoute() string {
	return api.server.URL + "/api/health"
}

func (api *TestingAPI) GetTenantsRoute() string {
	return api.server.URL + "/api/tenants"
}

func (api *TestingAPI) CreateTenant(tenant any) (*http.Response, error) {
	jsonTenant, _ := json.Marshal(tenant)
	return http.Post(api.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
}
