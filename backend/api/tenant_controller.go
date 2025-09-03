package api

import (
	"cochera/internal/formatting"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const TenantsBaseRoute string = "/tenants"

func (api *API) tenantHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getTenants(w, r)
	case http.MethodPost:
		api.createTenant(w, r)
	default:
		formatter := formatting.NewResponseFormatter(w)
		formatter.RespondMethodIsNotAllowed()
	}
}

func (api *API) getTenants(w http.ResponseWriter, _ *http.Request) {
	// _ := formatting.NewResponseFormatter(w)

	_, err := api.tenantService.GetAllTenants()
	if err != nil {
		// formatter.RespondCouldNotFindAnyTenants(err)
		return
	}

	// err = formatter.RespondTenantsGotSuccessfully(tenants)
	// if err != nil {
	// 	formatter.RespondCouldNotWriteResponse(err)
	// }
}

func (api *API) createTenant(w http.ResponseWriter, r *http.Request) {
	formatter := formatting.NewResponseFormatter(w)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		formatter.RespondCouldNotReadRequestBody()
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			formatter.RespondCouldNotCloseRequestBody()
		}
	}()

	tenant, err := api.tenantService.CreateTenant(requestBody)
	if err != nil {
		formatter.RespondCouldNotCreateTenant(err)
		return
	}

	err = formatter.RespondTenantWasCreatedSuccessfully(tenant)
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}

func (api *API) tenantByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getTenantByID(w, r)
	default:
		formatter := formatting.NewResponseFormatter(w)
		formatter.RespondMethodIsNotAllowed()
	}
}

func (api *API) getTenantByID(w http.ResponseWriter, r *http.Request) {
	formatter := formatting.NewResponseFormatter(w)

	path := strings.TrimPrefix(r.URL.Path, TenantsBaseRoute+"/")
	if path == "" {
		formatter.RespondTenantIDMustNotBeMissing()
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		formatter.RespondTenantIDMustBeAnInteger()
		return
	}

	tenant, err := api.tenantService.GetTenantByID(id)
	if err != nil {
		formatter.RespondCouldNotGetTenant(err)
		return
	}

	err = formatter.RespondTenantGotSuccessfully(tenant)
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}
