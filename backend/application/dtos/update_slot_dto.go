package dtos

import (
	vo "cochera/domain/value_objects"
)

type UpdateSlotDTO struct {
	Number   *vo.SlotNumber `json:"number"`
	TenantID *vo.EntityID   `json:"tenant_id"`
}
