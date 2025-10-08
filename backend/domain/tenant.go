package domain

type Tenant struct {
	ID         EntityID     `json:"id"`
	DNI        DNI          `json:"dni"`
	Name       Name         `json:"name"`
	LastName   Name         `json:"last_name"`
	Address    Address      `json:"address"`
	Phone      Phone        `json:"phone"`
	Email      EmailAddress `json:"email"`
	EntryMonth MonthOfYear  `json:"entry_month"`
}

func (tenant *Tenant) SetID(id int) {
	tenant.ID = EntityID(id)
}

func (tenant *Tenant) GetDNI() uint32 {
	return uint32(tenant.DNI)
}

func (tenant *Tenant) HasDNI(dni uint32) bool {
	return uint32(tenant.DNI) == dni
}

func (tenant *Tenant) GetName() string {
	return string(tenant.Name)
}

func (tenant *Tenant) HasName(name string) bool {
	return string(tenant.Name) == name
}

func (tenant *Tenant) GetLastName() string {
	return string(tenant.LastName)
}

func (tenant *Tenant) HasLastName(lastName string) bool {
	return string(tenant.LastName) == lastName
}

func (tenant *Tenant) GetAddress() string {
	return string(tenant.Address)
}

func (tenant *Tenant) GetPhone() string {
	return tenant.Phone.String()
}

func (tenant *Tenant) GetEmail() string {
	return string(tenant.Email)
}

func (tenant *Tenant) HasEmail(email string) bool {
	return string(tenant.Email) == email
}

func (tenant *Tenant) GetEntryMonth() string {
	return tenant.EntryMonth.String()
}

func NewTenant(id, dni, name, lastName, address, phone, email, entryMonth any) (*Tenant, error) {
	var tenant Tenant
	var err error

	tenant.ID, err = NewEntityID(id)
	if err != nil {
		return nil, err
	}

	tenant.DNI, err = NewDNI(dni)
	if err != nil {
		return nil, err
	}

	tenant.Name, err = NewName(name)
	if err != nil {
		return nil, err
	}

	tenant.LastName, err = NewName(lastName)
	if err != nil {
		return nil, err
	}

	tenant.Address, err = NewAddress(address)
	if err != nil {
		return nil, err
	}

	tenant.Phone, err = NewPhone(phone)
	if err != nil {
		return nil, err
	}

	tenant.Email, err = NewEmailAddress(email)
	if err != nil {
		return nil, err
	}

	tenant.EntryMonth, err = NewMonthOfYearFromString(entryMonth)
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

func (tenant *Tenant) Validate() error {
	if tenant.DNI == 0 {
		return ErrRequiredDNI
	}

	if tenant.Name == "" {
		return ErrRequiredName
	}

	if tenant.LastName == "" {
		return ErrRequiredLastName
	}

	if tenant.EntryMonth.Month == 0 && tenant.EntryMonth.Year == 0 {
		return ErrRequiredEntryMonth
	}

	return nil
}
