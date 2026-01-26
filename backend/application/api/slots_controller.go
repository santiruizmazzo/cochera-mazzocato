package api

import (
	"cochera/application/dtos"
	"cochera/application/formatting"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const SlotsBaseRoute string = "/slots"

func (api API) slotByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getSlotByID(w, r)
	case http.MethodPatch:
		api.updateSlotByID(w, r)
	default:
		formatter := formatting.NewResponseFormatter(w)
		formatter.RespondMethodIsNotAllowed()
	}
}

func (api API) getSlotByID(w http.ResponseWriter, r *http.Request) {
	formatter := formatting.NewResponseFormatter(w)

	path := strings.TrimPrefix(r.URL.Path, SlotsBaseRoute+"/")
	if path == "" {
		formatter.RespondResourceIDMustNotBeMissing()
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		formatter.RespondResourceIDMustBeAnInteger()
		return
	}

	slot, err := api.slotService.GetByID(id)
	if err != nil {
		formatter.RespondCouldNotGetSlot(err)
		return
	}

	err = formatter.RespondSlotGotSuccessfully(slot)
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}

func (api API) updateSlotByID(w http.ResponseWriter, r *http.Request) {
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

	path := strings.TrimPrefix(r.URL.Path, SlotsBaseRoute+"/")
	if path == "" {
		formatter.RespondResourceIDMustNotBeMissing()
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		formatter.RespondResourceIDMustBeAnInteger()
		return
	}

	var updateDTO dtos.UpdateSlotDTO

	if err := json.Unmarshal(requestBody, &updateDTO); err != nil {
		formatter.RespondCouldNotUpdateSlot(err)
		return
	}

	slot, err := api.slotService.UpdateSlot(id, updateDTO)
	if err != nil {
		formatter.RespondCouldNotUpdateSlot(err)
		return
	}

	err = formatter.RespondSlotUpdatedSuccessfully(slot)
	if err != nil {
		formatter.RespondCouldNotWriteResponse(err)
	}
}
