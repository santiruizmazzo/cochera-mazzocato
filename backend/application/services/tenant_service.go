package services

import (
	"cochera/domain"
)

type TenantService struct {
	repo domain.TenantRepository
}

func NewTenantService(repo domain.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (service *TenantService) GetTenantByID(id int) (*domain.Tenant, error) {
	return service.repo.GetTenantByID(id)
}

func (service *TenantService) GetAllTenants() ([]*domain.Tenant, error) {
	return service.repo.GetAllTenants()
}

func (service *TenantService) GetAllTenantsByName(name string) ([]*domain.Tenant, error) {
	return service.repo.GetAllTenantsByName(name)
}

func (service *TenantService) CreateTenant(jsonTenant []byte) (*domain.Tenant, error) {
	tenant, err := domain.NewTenantFromJSON(jsonTenant)
	if err != nil {
		return nil, err
	}

	dniAlreadyExists, err := service.repo.ExistsTenantWithDNI(tenant.DNI)
	if err != nil {
		return nil, err
	}

	if dniAlreadyExists {
		return nil, domain.ErrDuplicateDNI
	}

	emailAlreadyExists, err := service.repo.ExistsTenantWithEmail(tenant.Email)
	if err != nil {
		return nil, err
	}

	if emailAlreadyExists {
		return nil, domain.ErrDuplicateEmail
	}

	return service.repo.Save(tenant)
}
