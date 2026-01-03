package ent

import (
	vo "cochera/domain/value_objects"
	"errors"
)

type Tenant struct {
	ID         vo.EntityID     `json:"id"`
	DNI        vo.DNI          `json:"dni"`
	Name       vo.Name         `json:"name"`
	LastName   vo.Name         `json:"last_name"`
	Address    vo.Address      `json:"address"`
	Phone      vo.Phone        `json:"phone"`
	Email      vo.EmailAddress `json:"email"`
	EntryMonth vo.MonthOfYear  `json:"entry_month"`
}

func (tenant *Tenant) SetID(id int) {
	tenant.ID = vo.EntityID(id)
}

func (tenant Tenant) GetDNI() int {
	return int(tenant.DNI)
}

func (tenant *Tenant) SetDNI(dni vo.DNI) {
	tenant.DNI = dni
}

func (tenant Tenant) HasDNI(dni int) bool {
	return int(tenant.DNI) == dni
}

func (tenant Tenant) GetName() string {
	return string(tenant.Name)
}

func (tenant *Tenant) SetName(name vo.Name) {
	tenant.Name = name
}

func (tenant Tenant) HasName(name string) bool {
	return string(tenant.Name) == name
}

func (tenant Tenant) GetLastName() string {
	return string(tenant.LastName)
}

func (tenant Tenant) HasLastName(lastName string) bool {
	return string(tenant.LastName) == lastName
}

func (tenant Tenant) GetAddress() any {
	if tenant.Address.IsEmpty() {
		return nil
	}
	return string(tenant.Address)
}

func (tenant Tenant) GetPhone() any {
	if tenant.Phone.IsEmpty() {
		return nil
	}
	return tenant.Phone.String()
}

func (tenant Tenant) GetEmail() any {
	if tenant.Email.IsEmpty() {
		return nil
	}
	return string(tenant.Email)
}

func (tenant Tenant) HasEmail(email string) bool {
	return string(tenant.Email) == email
}

func (tenant Tenant) GetEntryMonth() string {
	return tenant.EntryMonth.String()
}

func (tenant *Tenant) SetEntryMonth(entryMonth vo.MonthOfYear) {
	tenant.EntryMonth = entryMonth
}

func NewTenant(id, dni, name, lastName, address, phone, email, entryMonth any) (*Tenant, error) {
	var tenant Tenant
	var err error

	if tenant.ID, err = vo.NewEntityID(id); err != nil {
		return nil, err
	}

	if tenant.DNI, err = vo.NewDNI(dni); err != nil {
		return nil, err
	}

	if tenant.Name, err = vo.NewName(name); err != nil {
		return nil, err
	}

	if tenant.LastName, err = vo.NewName(lastName); err != nil {
		return nil, err
	}

	if tenant.Address, err = vo.NewAddress(address); err != nil {
		return nil, err
	}

	if tenant.Phone, err = vo.NewPhone(phone); err != nil {
		return nil, err
	}

	if tenant.Email, err = vo.NewEmailAddress(email); err != nil {
		return nil, err
	}

	if tenant.EntryMonth, err = vo.NewMonthOfYear(entryMonth); err != nil {
		return nil, err
	}

	return &tenant, nil
}

var (
	ErrRequiredDNI        = errors.New("el DNI es obligatorio")
	ErrRequiredName       = errors.New("el nombre es obligatorio")
	ErrRequiredLastName   = errors.New("el apellido es obligatorio")
	ErrRequiredEntryMonth = errors.New("el mes de ingreso es obligatorio")
)

func (tenant Tenant) Validate() error {
	switch true {
	case tenant.DNI == 0:
		return ErrRequiredDNI
	case tenant.Name == "":
		return ErrRequiredName
	case tenant.LastName == "":
		return ErrRequiredLastName
	case tenant.EntryMonth.Month == 0 && tenant.EntryMonth.Year == 0:
		return ErrRequiredEntryMonth
	default:
		return nil
	}
}
