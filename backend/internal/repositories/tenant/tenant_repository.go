package tenantrepo

import "cochera/internal/domain/tenant"

type TenantRepository interface {
	GetTenantByID(id int) (*tenant.Tenant, error)
	GetAllTenants() ([]*tenant.Tenant, error)
	Save(tenant *tenant.Tenant) (*tenant.Tenant, error)
	ExistsTenantWithDNI(dni uint32) (bool, error)
	ExistsTenantWithEmail(email string) (bool, error)
}
