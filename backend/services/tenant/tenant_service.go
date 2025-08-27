package tenantservice

import (
	"cochera/domain/tenant"
	tenantrepo "cochera/repositories/tenant"
)

type TenantService struct {
	repo tenantrepo.TenantRepository
}

func NewTenantService(repo tenantrepo.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (service *TenantService) CreateTenant(jsonTenant []byte) (*tenant.Tenant, error) {
	tenant, err := tenant.NewTenantFromJSON(jsonTenant)
	if err != nil {
		return nil, err
	}

	return service.repo.Save(tenant)
}
