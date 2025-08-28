package api

import (
	"cochera/internal/domain/tenant"
	"net/http"
)

const TenantsBaseRoute string = "/tenants"

func (api *API) createTenant(w http.ResponseWriter, r *http.Request) {
	var requestBody []byte
	if _, err := r.Body.Read(requestBody); err != nil {
		http.Error(w, "Error al leer request body", http.StatusInternalServerError)
	}

	_, err := tenant.NewTenantFromJSON(requestBody)
	if err != nil {
		http.Error(w, "Inquilino invalido", http.StatusBadRequest)
	}

}
