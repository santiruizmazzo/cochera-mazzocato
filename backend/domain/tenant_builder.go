package domain

import vo "cochera/domain/value_objects"

type TenantBuilder struct {
	tenant *Tenant
}

func NewTenantBuilder() *TenantBuilder {
	tenant := Tenant{
		ID:         1,
		DNI:        12345678,
		Name:       "Huang",
		LastName:   "Lee",
		Address:    "22 Beer Heights St.",
		Phone:      vo.Phone{CountryCode: "54", LineNumber: "9815111526"},
		Email:      "huang@lee.com",
		EntryMonth: vo.MonthOfYear{Month: 1, Year: 2025},
	}
	return &TenantBuilder{tenant: &tenant}
}

func (builder *TenantBuilder) WithID(id int) *TenantBuilder {
	builder.tenant.ID = vo.EntityID(id)
	return builder
}

func (builder *TenantBuilder) WithDNI(dni int) *TenantBuilder {
	builder.tenant.DNI = vo.DNI(dni)
	return builder
}

func (builder *TenantBuilder) WithName(name string) *TenantBuilder {
	builder.tenant.Name = vo.Name(name)
	return builder
}

func (builder *TenantBuilder) WithLastName(lastName string) *TenantBuilder {
	builder.tenant.LastName = vo.Name(lastName)
	return builder
}

func (builder *TenantBuilder) WithAddress(address string) *TenantBuilder {
	builder.tenant.Address = vo.Address(address)
	return builder
}

func (builder *TenantBuilder) WithPhone(phone string) *TenantBuilder {
	builder.tenant.Phone, _ = vo.NewPhone(phone)
	return builder
}

func (builder *TenantBuilder) WithEmail(email string) *TenantBuilder {
	builder.tenant.Email = vo.EmailAddress(email)
	return builder
}

func (builder *TenantBuilder) Build() *Tenant {
	return builder.tenant
}
