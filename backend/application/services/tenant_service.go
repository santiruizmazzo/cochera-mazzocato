package tenantservice

import (
	"cochera/domain"
	myerrors "cochera/domain/errors"
	"cochera/domain/tenant"
)

type TenantService struct {
	repo domain.TenantRepository
}

func NewTenantService(repo domain.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (service *TenantService) GetTenantByID(id int) (*tenant.Tenant, error) {
	return service.repo.GetTenantByID(id)
}

func (service *TenantService) GetAllTenants() ([]*tenant.Tenant, error) {
	return service.repo.GetAllTenants()
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
