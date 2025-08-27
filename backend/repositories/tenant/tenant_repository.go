package tenantrepo

import "cochera/domain/tenant"

type TenantRepository interface {
	Save(tenant *tenant.Tenant) (*tenant.Tenant, error)
}
