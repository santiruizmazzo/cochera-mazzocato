package api

import (
	myerrors "cochera/internal/errors"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const TenantsBaseRoute string = "/tenants"

func (api *API) createTenant(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			http.Error(w, "Could not close request body", http.StatusInternalServerError)
		}
	}()

	tenant, err := api.tenantService.CreateTenant(requestBody)
	if err != nil {
		statusCode := http.StatusBadRequest
		if errors.Is(err, myerrors.ErrDuplicateDNI) {
			statusCode = http.StatusConflict
		}
		responseBody := fmt.Sprintf(`{"detail":"%s"}`, err.Error())
		http.Error(w, responseBody, statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(tenant)
	if err != nil {
		responseBody := fmt.Sprintf(`{"detail":"%s"}`, err.Error())
		http.Error(w, responseBody, http.StatusInternalServerError)
		return
	}
}
