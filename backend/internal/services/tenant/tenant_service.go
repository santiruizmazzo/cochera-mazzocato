package tenantservice

import (
	"cochera/internal/domain/tenant"
	myerrors "cochera/internal/errors"
	tenantrepo "cochera/internal/repositories/tenant"
)

type TenantService struct {
	repo tenantrepo.TenantRepository
}

func NewTenantService(repo tenantrepo.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (service *TenantService) GetTenantByID(id int) (*tenant.Tenant, error) {
	return service.repo.GetTenantByID(id)
}

func (service *TenantService) CreateTenant(jsonTenant []byte) (*tenant.Tenant, error) {
	tenant, err := tenant.NewTenantFromJSON(jsonTenant)
	if err != nil {
		return nil, err
	}

	dniAlreadyExists, err := service.repo.ExistsTenantWithDNI(tenant.DNI)
	if err != nil {
		return nil, err
	}

	if dniAlreadyExists {
		return nil, myerrors.ErrDuplicateDNI
	}

	emailAlreadyExists, err := service.repo.ExistsTenantWithEmail(tenant.Email)
	if err != nil {
		return nil, err
	}

	if emailAlreadyExists {
		return nil, myerrors.ErrDuplicateEmail
	}

	return service.repo.Save(tenant)
}
