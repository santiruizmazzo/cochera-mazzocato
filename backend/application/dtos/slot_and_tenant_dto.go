package dtos

import (
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
)

type SlotAndTenantDTO struct {
	ID     vo.EntityID      `json:"id"`
	Number vo.SlotNumber    `json:"number"`
	Tenant *TenantInSlotDTO `json:"tenant"`
}

func NewSlotAndTenantDTO(slot *ent.Slot, tenant *ent.Tenant) *SlotAndTenantDTO {
	return &SlotAndTenantDTO{
		ID:     slot.ID,
		Number: slot.Number,
		Tenant: NewTenantInSlotDTO(tenant),
	}
}
