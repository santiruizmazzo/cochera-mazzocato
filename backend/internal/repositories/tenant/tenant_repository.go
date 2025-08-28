package tenantrepo

import "cochera/internal/domain/tenant"

type TenantRepository interface {
	Save(tenant *tenant.Tenant) (*tenant.Tenant, error)
}
