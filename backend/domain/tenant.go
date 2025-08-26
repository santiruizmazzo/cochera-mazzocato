package domain

import (
	"encoding/json"
)

type Tenant struct {
	ID         uint32
	DNI        uint32      `json:"dni"`
	Name       string      `json:"name"`
	LastName   string      `json:"last_name"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone"`
	Email      string      `json:"email"`
	EntryMonth MonthOfYear `json:"entry_month"`
}

func NewTenantFromJSON(jsonBytes []byte) (*Tenant, error) {
	var tenant Tenant

	if err := json.Unmarshal(jsonBytes, &tenant); err != nil {
		return nil, err
	}
	return &tenant, nil
}
