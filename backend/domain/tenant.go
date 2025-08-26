package domain

type Tenant struct {
	id         uint32
	dni        uint32
	name       string
	lastName   string
	address    string
	phone      string
	email      string
	entryMonth MonthOfYear
}

func NewTenantFromJSON(jsonBytes []byte) (*Tenant, error) {
	panic("unimplemented")
}
