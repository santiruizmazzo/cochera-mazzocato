package dtos

import (
	"cochera/domain"
	vo "cochera/domain/value_objects"
)

type TenantListDTO struct {
	ID         vo.EntityID    `json:"id"`
	Name       vo.Name        `json:"name"`
	LastName   vo.Name        `json:"last_name"`
	EntryMonth vo.MonthOfYear `json:"entry_month"`
}

func NewTenantListDTO(tenant *domain.Tenant) *TenantListDTO {
	return &TenantListDTO{
		ID:         tenant.ID,
		Name:       tenant.Name,
		LastName:   tenant.LastName,
		EntryMonth: tenant.EntryMonth,
	}
}
