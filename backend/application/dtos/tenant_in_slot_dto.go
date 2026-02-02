package dtos

import (
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
)

type TenantInSlotDTO struct {
	ID       vo.EntityID `json:"id"`
	Name     vo.Name     `json:"name"`
	LastName vo.Name     `json:"last_name"`
}

func NewTenantInSlotDTO(tenant *ent.Tenant) *TenantInSlotDTO {
	if tenant == nil {
		return nil
	}

	return &TenantInSlotDTO{
		ID:       tenant.ID,
		Name:     tenant.Name,
		LastName: tenant.LastName,
	}
}
