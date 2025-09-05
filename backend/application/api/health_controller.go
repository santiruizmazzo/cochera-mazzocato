package api

import (
	"cochera/application"
	"cochera/application/formatting"
	"net/http"
)

const HealthRoute string = "/health"

func (api *API) getHealthStatus(w http.ResponseWriter, r *http.Request) {
	formatter := formatting.NewResponseFormatter(w)

	err := formatter.RespondCurrentHealthStatus(application.CurrentVersion())
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}
