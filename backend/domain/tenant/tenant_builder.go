package tenant

import "cochera/domain/calendar"

type TenantBuilder struct {
	tenant *Tenant
}

func NewTenantBuilder() *TenantBuilder {
	tenant := Tenant{
		DNI:        12345678,
		Name:       "Huang",
		LastName:   "Lee",
		Address:    "22 Beer Heights St.",
		Phone:      "+98151526",
		Email:      "huang@lee.com",
		EntryMonth: calendar.NewMonthOfYear(1, 2025),
	}
	return &TenantBuilder{tenant: &tenant}
}

func (builder *TenantBuilder) WithID(id int) *TenantBuilder {
	builder.tenant.ID = uint32(id)
	return builder
}

func (builder *TenantBuilder) WithDNI(dni int) *TenantBuilder {
	builder.tenant.DNI = uint32(dni)
	return builder
}

func (builder *TenantBuilder) WithName(name string) *TenantBuilder {
	builder.tenant.Name = name
	return builder
}

func (builder *TenantBuilder) WithLastName(lastName string) *TenantBuilder {
	builder.tenant.LastName = lastName
	return builder
}

func (builder *TenantBuilder) WithAddress(address string) *TenantBuilder {
	builder.tenant.Address = address
	return builder
}

func (builder *TenantBuilder) WithPhone(phone string) *TenantBuilder {
	builder.tenant.Phone = phone
	return builder
}

func (builder *TenantBuilder) WithEmail(email string) *TenantBuilder {
	builder.tenant.Email = email
	return builder
}

func (builder *TenantBuilder) Build() *Tenant {
	return builder.tenant
}
