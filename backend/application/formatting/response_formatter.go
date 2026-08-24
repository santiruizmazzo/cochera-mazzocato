package formatting

import (
	"cochera/application/dtos"
	"cochera/application/services"
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
	"cochera/infra"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ResponseFormatter is a utility type for controlling API's response formats in a centralized and homogeneous way.
type ResponseFormatter struct {
	w http.ResponseWriter
}

func NewResponseFormatter(w http.ResponseWriter) *ResponseFormatter {
	return &ResponseFormatter{w: w}
}

func (formatter ResponseFormatter) CreateErrorBodyWith(errorDetail string) string {
	errorResponseScheme := `{"detail":"%s"}`
	return fmt.Sprintf(errorResponseScheme, errorDetail)
}

func (formatter ResponseFormatter) RespondMethodIsNotAllowed() {
	responseBody := formatter.CreateErrorBodyWith("method not allowed")
	http.Error(formatter.w, responseBody, http.StatusMethodNotAllowed)
}

func (formatter ResponseFormatter) RespondCouldNotReadRequestBody() {
	responseBody := formatter.CreateErrorBodyWith("could not read request body")
	http.Error(formatter.w, responseBody, http.StatusBadRequest)
}

func (formatter ResponseFormatter) RespondCouldNotCloseRequestBody() {
	responseBody := formatter.CreateErrorBodyWith("could not close request body")
	http.Error(formatter.w, responseBody, http.StatusInternalServerError)
}

func (formatter ResponseFormatter) RespondCouldNotCreateTenant(creationError error) {
	statusCode := http.StatusBadRequest

	if errors.Is(creationError, services.ErrDuplicateDNI) || errors.Is(creationError, services.ErrDuplicateEmail) {
		statusCode = http.StatusConflict
	}

	responseBody := formatter.CreateErrorBodyWith(creationError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter ResponseFormatter) RespondTenantWasCreatedSuccessfully(tenant *ent.Tenant) error {
	formatter.w.Header().Set("Content-Type", "application/json")
	formatter.w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(formatter.w).Encode(tenant)
}

func (formatter ResponseFormatter) RespondCouldNotWriteResponse(responseError error) {
	responseBody := formatter.CreateErrorBodyWith(responseError.Error())
	http.Error(formatter.w, responseBody, http.StatusInternalServerError)
}

func (formatter ResponseFormatter) RespondResourceIDMustNotBeMissing() {
	responseBody := formatter.CreateErrorBodyWith("id is required")
	http.Error(formatter.w, responseBody, http.StatusBadRequest)
}

func (formatter ResponseFormatter) RespondResourceIDMustBeAnInteger() {
	responseBody := formatter.CreateErrorBodyWith("id must be an integer")
	http.Error(formatter.w, responseBody, http.StatusBadRequest)
}

func (formatter ResponseFormatter) RespondCouldNotGetTenant(retrievingError error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(retrievingError, infra.ErrTenantNotFound) {
		statusCode = http.StatusNotFound
	}

	responseBody := formatter.CreateErrorBodyWith(retrievingError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter ResponseFormatter) RespondTenantGotSuccessfully(tenant *ent.Tenant) error {
	formatter.w.Header().Set("Content-Type", "application/json")
	formatter.w.WriteHeader(http.StatusOK)
	return json.NewEncoder(formatter.w).Encode(tenant)
}

func (formatter ResponseFormatter) RespondCurrentHealthStatus(currentVersion string) error {
	response := map[string]string{"status": "live", "version": currentVersion}
	formatter.w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(formatter.w).Encode(response)
}

func (formatter ResponseFormatter) RespondCouldNotFindAnyResources(retrievingError error) {
	statusCode := http.StatusNotFound
	responseBody := formatter.CreateErrorBodyWith(retrievingError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter ResponseFormatter) RespondTenantsGotSuccessfully(tenants []*ent.Tenant) error {
	responseData := make([]*dtos.TenantListDTO, len(tenants))
	for i, tenant := range tenants {
		responseData[i] = dtos.NewTenantListDTO(tenant)
	}

	response := map[string]any{"data": responseData}
	formatter.w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(formatter.w).Encode(response)
}

func (formatter *ResponseFormatter) RespondCouldNotUpdateTenant(updateError error) {
	statusCode := http.StatusBadRequest

	if errors.Is(updateError, infra.ErrTenantNotFound) {
		statusCode = http.StatusNotFound
	}

	responseBody := formatter.CreateErrorBodyWith(updateError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter *ResponseFormatter) RespondTenantUpdatedSuccessfully(tenant *ent.Tenant) error {
	formatter.w.Header().Set("Content-Type", "application/json")
	formatter.w.WriteHeader(http.StatusOK)
	return json.NewEncoder(formatter.w).Encode(tenant)
}

func (formatter *ResponseFormatter) RespondCouldNotUpdateSlot(updateError error) {
	statusCode := http.StatusBadRequest

	if errors.Is(updateError, infra.ErrSlotNotFound) || errors.Is(updateError, infra.ErrTenantNotFound) {
		statusCode = http.StatusNotFound
	}

	responseBody := formatter.CreateErrorBodyWith(updateError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter *ResponseFormatter) RespondSlotUpdatedSuccessfully(slot *ent.Slot) error {
	formatter.w.Header().Set("Content-Type", "application/json")
	formatter.w.WriteHeader(http.StatusOK)
	return json.NewEncoder(formatter.w).Encode(slot)
}

func (formatter *ResponseFormatter) RespondCouldNotGetSlot(retrievingError error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(retrievingError, infra.ErrSlotNotFound) {
		statusCode = http.StatusNotFound
	}

	responseBody := formatter.CreateErrorBodyWith(retrievingError.Error())
	http.Error(formatter.w, responseBody, statusCode)
}

func (formatter ResponseFormatter) RespondSlotGotSuccessfully(slot *ent.Slot) error {
	formatter.w.Header().Set("Content-Type", "application/json")
	formatter.w.WriteHeader(http.StatusOK)
	return json.NewEncoder(formatter.w).Encode(slot)
}

func (formatter ResponseFormatter) RespondSlotsGotSuccessfully(slots []*ent.Slot, tenants []*ent.Tenant) error {
	mapOfTenants := map[vo.EntityID]*ent.Tenant{}
	for _, tenant := range tenants {
		mapOfTenants[tenant.ID] = tenant
	}

	listOfSlots := make([]dtos.SlotAndTenantDTO, 0, len(slots))
	for _, slot := range slots {
		tenant := mapOfTenants[slot.TenantID]
		listOfSlots = append(listOfSlots, *dtos.NewSlotAndTenantDTO(slot, tenant))
	}

	response := map[string]any{"data": listOfSlots}

	formatter.w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(formatter.w).Encode(response)
}
