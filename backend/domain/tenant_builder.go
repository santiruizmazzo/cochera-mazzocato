package domain

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
		Phone:      Phone{CountryCode: "54", LineNumber: "9815111526"},
		Email:      "huang@lee.com",
		EntryMonth: MonthOfYear{Month: 1, Year: 2025},
	}
	return &TenantBuilder{tenant: &tenant}
}

func (builder *TenantBuilder) WithID(id int) *TenantBuilder {
	builder.tenant.ID = EntityID(id)
	return builder
}

func (builder *TenantBuilder) WithDNI(dni int) *TenantBuilder {
	builder.tenant.DNI = DNI(dni)
	return builder
}

func (builder *TenantBuilder) WithName(name string) *TenantBuilder {
	builder.tenant.Name = Name(name)
	return builder
}

func (builder *TenantBuilder) WithLastName(lastName string) *TenantBuilder {
	builder.tenant.LastName = Name(lastName)
	return builder
}

func (builder *TenantBuilder) WithAddress(address string) *TenantBuilder {
	builder.tenant.Address = Address(address)
	return builder
}

func (builder *TenantBuilder) WithPhone(phone string) *TenantBuilder {
	builder.tenant.Phone, _ = NewPhone(phone)
	return builder
}

func (builder *TenantBuilder) WithEmail(email string) *TenantBuilder {
	builder.tenant.Email = EmailAddress(email)
	return builder
}

func (builder *TenantBuilder) Build() *Tenant {
	return builder.tenant
}
