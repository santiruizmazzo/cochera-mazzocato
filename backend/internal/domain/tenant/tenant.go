package tenant

import (
	"cochera/internal/domain/time"
	"encoding/json"
	"errors"
	"fmt"
)

type Tenant struct {
	ID         uint32
	DNI        uint32           `json:"dni"`
	Name       string           `json:"name"`
	LastName   string           `json:"last_name"`
	Address    string           `json:"address"`
	Phone      string           `json:"phone"`
	Email      string           `json:"email"`
	EntryMonth time.MonthOfYear `json:"entry_month"`
}

func NewTenantFromJSON(jsonBytes []byte) (*Tenant, error) {
	var tenant Tenant

	if err := json.Unmarshal(jsonBytes, &tenant); err != nil {
		return nil, fmt.Errorf("could not create tenant from json: %w", err)
	}

	if err := tenant.ValidateAttributes(); err != nil {
		return nil, errors.New("required attributes: dni, name, last_name or entry_month")
	}

	return &tenant, nil
}

func (tenant *Tenant) ValidateAttributes() error {
	if tenant.DNI == 0 {
		return errors.New("dni is required")
	}

	if tenant.Name == "" {
		return errors.New("name is required")
	}

	if tenant.LastName == "" {
		return errors.New("last_name is required")
	}

	if tenant.EntryMonth.String() == "00-0000" {
		return errors.New("entry_month is required")
	}

	return nil
}
