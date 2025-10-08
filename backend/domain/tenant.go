package domain

type Tenant struct {
	// ID         uint32      `json:"id"`
	// DNI        uint32      `json:"dni"`
	// Name       string      `json:"name"`
	// LastName   string      `json:"last_name"`
	// Address    string      `json:"address"`
	// Phone      string      `json:"phone"`
	// Email      string      `json:"email"`
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

// func NewTenant(dni, name, lastName, address, phone, email, entryMonth any) (*Tenant, error) {
// 	validDNI, err := validateDNI(dni)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validName, err := validateName(name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validLastName, err := validateLastName(lastName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validAddress, err := validateAddress(address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validPhone, err := validatePhone(phone)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validEmail, err := validateEmail(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	validEntryMonth, err := validateEntryMonth(entryMonth)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Tenant{
// 		DNI:        validDNI,
// 		Name:       validName,
// 		LastName:   validLastName,
// 		Address:    validAddress,
// 		Phone:      validPhone,
// 		Email:      validEmail,
// 		EntryMonth: validEntryMonth,
// 	}, nil
// }

// func NewTenantFromJSON(jsonBytes []byte) (*Tenant, error) {
// 	var tenantMap map[string]any

// 	if err := json.Unmarshal(jsonBytes, &tenantMap); err != nil {
// 		return nil, fmt.Errorf("couldn't read json: %w", err)
// 	}

// 	return buildValidTenant(tenantMap)
// }

// func buildValidTenant(tenantMap map[string]any) (*Tenant, error) {
// 	dni, err := extractAndValidateDNI(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	name, err := extractAndValidateName(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lastName, err := extractAndValidateLastName(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	address, err := extractAndValidateAddress(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	phone, err := extractAndValidatePhone(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	email, err := extractAndValidateEmail(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	entryMonth, err := extractAndValidateEntryMonth(tenantMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return NewTenant(dni, name, lastName, address, phone, email, entryMonth)
// }

// func (tenant Tenant) MarshalJSON() ([]byte, error) {
// 	result := map[string]any{
// 		"id":          tenant.ID,
// 		"dni":         tenant.DNI,
// 		"name":        tenant.Name,
// 		"last_name":   tenant.LastName,
// 		"entry_month": tenant.EntryMonth.String(),
// 	}

// 	if tenant.Address != "" {
// 		result["address"] = tenant.Address
// 	} else {
// 		result["address"] = nil
// 	}

// 	if tenant.Phone != "" {
// 		result["phone"] = tenant.Phone
// 	} else {
// 		result["phone"] = nil
// 	}

// 	if tenant.Email != "" {
// 		result["email"] = tenant.Email
// 	} else {
// 		result["email"] = nil
// 	}

// 	return json.Marshal(result)
// }
