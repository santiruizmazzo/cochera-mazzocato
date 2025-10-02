package dtos

import "cochera/domain"

type TenantListDTO struct {
	ID         uint32             `json:"id"`
	Name       string             `json:"name"`
	LastName   string             `json:"last_name"`
	EntryMonth domain.MonthOfYear `json:"entry_month"`
}

func NewTenantListDTO(tenant *domain.Tenant) *TenantListDTO {
	return &TenantListDTO{
		ID:         tenant.ID,
		Name:       tenant.Name,
		LastName:   tenant.LastName,
		EntryMonth: tenant.EntryMonth,
	}
}
