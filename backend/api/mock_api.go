package api

import (
	"cochera/handlers"
	"net/http"
	"net/http/httptest"
)

type MockAPI struct {
	server *httptest.Server
	router *http.ServeMux
}

func NewMock() *MockAPI {
	router := setupEndpointsRouter()
	return &MockAPI{router: router}
}

func (api *MockAPI) Run() {
	api.server = httptest.NewServer(api.router)
}

func (api *MockAPI) Stop() {
	api.server.Close()
}

func (api *MockAPI) GetFullHealthRoute() string {
	return api.server.URL + handlers.HealthRoute
}
