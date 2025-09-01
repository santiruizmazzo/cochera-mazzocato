package api

import (
	myerrors "cochera/internal/errors"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const TenantsBaseRoute string = "/tenants"

func (api *API) tenantHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		api.createTenant(w, r)
	default:
		http.Error(w, `{"detail":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

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
		if errors.Is(err, myerrors.ErrDuplicateDNI) || errors.Is(err, myerrors.ErrDuplicateEmail) {
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

func (api *API) tenantByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getTenantByID(w, r)
	default:
		http.Error(w, `{"detail":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (api *API) getTenantByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, TenantsBaseRoute+"/")
	if path == "" {
		http.Error(w, `{"detail":"tenant id required"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, `{"detail":"invalid tenant id format"}`, http.StatusBadRequest)
		return
	}

	tenant, err := api.tenantService.GetTenantByID(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, myerrors.ErrTenantNotFound) {
			statusCode = http.StatusNotFound
		}
		responseBody := fmt.Sprintf(`{"detail":"%s"}`, err.Error())
		http.Error(w, responseBody, statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tenant)
	if err != nil {
		responseBody := fmt.Sprintf(`{"detail":"%s"}`, err.Error())
		http.Error(w, responseBody, http.StatusInternalServerError)
		return
	}
}
