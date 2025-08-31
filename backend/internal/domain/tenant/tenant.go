package tenant

import (
	"cochera/internal/domain/calendar"
	"encoding/json"
	"fmt"
)

type Tenant struct {
	ID         uint32               `json:"id,omitempty"`
	DNI        uint32               `json:"dni,omitempty"`
	Name       string               `json:"name,omitempty"`
	LastName   string               `json:"last_name,omitempty"`
	Address    string               `json:"address,omitempty"`
	Phone      string               `json:"phone,omitempty"`
	Email      string               `json:"email,omitempty"`
	EntryMonth calendar.MonthOfYear `json:"entry_month"`
}

func NewTenant(dni, name, lastName, address, phone, email, entryMonth any) (*Tenant, error) {
	validDNI, err := validateDNI(dni)
	if err != nil {
		return nil, err
	}

	validName, err := validateName(name)
	if err != nil {
		return nil, err
	}

	validLastName, err := validateLastName(lastName)
	if err != nil {
		return nil, err
	}

	validAddress, err := validateAddress(address)
	if err != nil {
		return nil, err
	}

	validPhone, err := validatePhone(phone)
	if err != nil {
		return nil, err
	}

	validEmail, err := validateEmail(email)
	if err != nil {
		return nil, err
	}

	validEntryMonth, err := validateEntryMonth(entryMonth)
	if err != nil {
		return nil, err
	}

	return &Tenant{
		DNI:        validDNI,
		Name:       validName,
		LastName:   validLastName,
		Address:    validAddress,
		Phone:      validPhone,
		Email:      validEmail,
		EntryMonth: validEntryMonth,
	}, nil
}

func NewTenantFromJSON(jsonBytes []byte) (*Tenant, error) {
	var tenantMap map[string]any

	if err := json.Unmarshal(jsonBytes, &tenantMap); err != nil {
		return nil, fmt.Errorf("couldn't read json: %w", err)
	}

	return buildValidTenant(tenantMap)
}

func buildValidTenant(tenantMap map[string]any) (*Tenant, error) {
	dni, err := extractAndValidateDNI(tenantMap)
	if err != nil {
		return nil, err
	}

	name, err := extractAndValidateName(tenantMap)
	if err != nil {
		return nil, err
	}

	lastName, err := extractAndValidateLastName(tenantMap)
	if err != nil {
		return nil, err
	}

	address, err := extractAndValidateAddress(tenantMap)
	if err != nil {
		return nil, err
	}

	phone, err := extractAndValidatePhone(tenantMap)
	if err != nil {
		return nil, err
	}

	email, err := extractAndValidateEmail(tenantMap)
	if err != nil {
		return nil, err
	}

	entryMonth, err := extractAndValidateEntryMonth(tenantMap)
	if err != nil {
		return nil, err
	}

	return NewTenant(dni, name, lastName, address, phone, email, entryMonth)
}
