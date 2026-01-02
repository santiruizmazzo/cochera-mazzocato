package dtos

import (
	vo "cochera/domain/value_objects"
)

type InputTenantDTO struct {
	DNI        vo.DNI          `json:"dni"`
	Name       vo.Name         `json:"name"`
	LastName   vo.Name         `json:"last_name"`
	Address    vo.Address      `json:"address"`
	Phone      vo.Phone        `json:"phone"`
	Email      vo.EmailAddress `json:"email"`
	EntryMonth vo.MonthOfYear  `json:"entry_month"`
}
