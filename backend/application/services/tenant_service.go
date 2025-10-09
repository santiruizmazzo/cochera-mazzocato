package services

import (
	"cochera/domain"
	ent "cochera/domain/entities"
	"encoding/json"
	"errors"
)

type TenantService struct {
	repo domain.TenantRepository
}

var (
	ErrDuplicateDNI   = errors.New("el DNI ya existe")
	ErrDuplicateEmail = errors.New("el email ya está en uso")
)

func NewTenantService(repo domain.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (service TenantService) GetByID(id int) (*ent.Tenant, error) {
	return service.repo.GetByID(id)
}

func (service TenantService) GetAllTenants() ([]*ent.Tenant, error) {
	return service.repo.GetAllTenants()
}

func (service TenantService) GetAllTenantsByName(name string) ([]*ent.Tenant, error) {
	return service.repo.GetAllTenantsByName(name)
}

func (service TenantService) GetAllTenantsByLastName(lastName string) ([]*ent.Tenant, error) {
	return service.repo.GetAllTenantsByLastName(lastName)
}

func (service TenantService) CreateTenant(jsonTenant []byte) (*ent.Tenant, error) {
	var tenant ent.Tenant

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
		return nil, ErrDuplicateDNI
	}

	emailAlreadyExists, err := service.repo.ExistsTenantWithEmail(tenant.Email.String())
	if err != nil {
		return nil, err
	}

	if emailAlreadyExists {
		return nil, ErrDuplicateEmail
	}

	return service.repo.Save(&tenant)
}
