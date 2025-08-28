package api

import (
	"cochera/internal/domain/tenant"
	"fmt"
	"io"
	"net/http"
)

const TenantsBaseRoute string = "/tenants"

func (api *API) createTenant(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer request body", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			http.Error(w, "Error al cerrar request body", http.StatusInternalServerError)
		}
	}()

	_, err = tenant.NewTenantFromJSON(requestBody)
	if err != nil {
		responseBody := fmt.Sprintf(`{"detail":"%s"}`, err.Error())
		http.Error(w, responseBody, http.StatusBadRequest)
		return
	}

}
