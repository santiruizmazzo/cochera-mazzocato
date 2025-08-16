package api

import (
	"net/http"
	"net/http/httptest"
)

type MockAPI struct {
	server *httptest.Server
	router *http.ServeMux
}

func NewMock() *MockAPI {
	api := &API{server: nil, db: nil}
	return &MockAPI{router: api.setupRouter()}
}

func (api *MockAPI) Run() {
	api.server = httptest.NewServer(api.router)
}

func (api *MockAPI) Stop() {
	api.server.Close()
}

func (api *MockAPI) GetHealthFullRoute() string {
	return api.server.URL + "/health"
}
