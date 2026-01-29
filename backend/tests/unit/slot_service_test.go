package unit

import (
	"cochera/application/dtos"
	"cochera/application/services"
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
	"cochera/infra"
	"encoding/json"
	"testing"
)

func TestSlotService_UpdateSlot_Successfully(t *testing.T) {
	existingSlot := ent.Slot{
		ID:       1,
		Number:   1,
		TenantID: 1,
	}

	existingTenant := ent.Tenant{
		ID: 1,
	}

	mockSlotRepo := &infra.InMemorySlotRepository{Slots: map[int]*ent.Slot{1: &existingSlot}}
	mockTenantRepo := &infra.InMemoryTenantRepository{Tenants: map[int]*ent.Tenant{6: &existingTenant}}

	service := services.NewSlotService(mockSlotRepo, mockTenantRepo)

	requestBody, _ := json.Marshal(map[string]any{
		"tenant_id": 6,
	})

	var updateDTO dtos.UpdateSlotDTO
	_ = json.Unmarshal(requestBody, &updateDTO)

	updatedSlot, err := service.UpdateSlot(1, updateDTO)

	if err != nil {
		t.Fatal("UpdateSlot should not fail: ", err)
	}

	expectedSlot := ent.Slot{ID: vo.EntityID(1), Number: 1, TenantID: 6}

	if *updatedSlot != expectedSlot {
		t.Fatal("Local Slot different from the updated one: ", updatedSlot)
	}
}
