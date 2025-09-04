package api

import (
	"cochera/application/formatting"
	"cochera/application/version"
	"net/http"
)

const HealthRoute string = "/health"

func (api *API) getHealthStatus(w http.ResponseWriter, r *http.Request) {
	formatter := formatting.NewResponseFormatter(w)

	err := formatter.RespondCurrentHealthStatus(version.Current())
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}
