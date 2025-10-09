package services

import (
	"cochera/domain"
	"encoding/json"
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

func (service *TenantService) GetAllTenantsByLastName(lastName string) ([]*domain.Tenant, error) {
	return service.repo.GetAllTenantsByLastName(lastName)
}

func (service *TenantService) CreateTenant(jsonTenant []byte) (*domain.Tenant, error) {
	var tenant domain.Tenant

	if err := json.Unmarshal(jsonTenant, &tenant); err != nil {
		return nil, err
	}

	if err := tenant.Validate(); err != nil {
		return nil, err
	}

	dniAlreadyExists, err := service.repo.ExistsTenantWithDNI(tenant.GetDNI())
	if err != nil {
		return nil, err
	}

	if dniAlreadyExists {
		return nil, domain.ErrDuplicateDNI
	}

	emailAlreadyExists, err := service.repo.ExistsTenantWithEmail(tenant.Email.String())
	if err != nil {
		return nil, err
	}

	if emailAlreadyExists {
		return nil, domain.ErrDuplicateEmail
	}

	return service.repo.Save(&tenant)
}
