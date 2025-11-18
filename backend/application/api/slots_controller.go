package api

import (
	"cochera/application/formatting"
	"net/http"
)

const SlotsBaseRoute string = "/slots"

func (api API) slotByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		api.allocateSlot(w, r)
	default:
		formatter := formatting.NewResponseFormatter(w)
		formatter.RespondMethodIsNotAllowed()
	}
}

func (api API) allocateSlot(w http.ResponseWriter, r *http.Request) {
	// formatter := formatting.NewResponseFormatter(w)

	// err := formatter.RespondCurrentHealthStatus(application.CurrentVersion())
	// if err != nil {
	// 	formatter.RespondCouldNotWriteResponse(err)
	// }
}
