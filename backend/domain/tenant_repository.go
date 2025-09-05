package domain

type TenantRepository interface {
	GetTenantByID(id int) (*Tenant, error)
	GetAllTenants() ([]*Tenant, error)
	Save(tenant *Tenant) (*Tenant, error)
	ExistsTenantWithDNI(dni uint32) (bool, error)
	ExistsTenantWithEmail(email string) (bool, error)
}
