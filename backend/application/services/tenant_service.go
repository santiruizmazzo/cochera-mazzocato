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

func (service TenantService) GetAll() ([]*ent.Tenant, error) {
	return service.repo.GetAll()
}

func (service TenantService) GetAllWithName(name string) ([]*ent.Tenant, error) {
	return service.repo.GetAllWithName(name)
}

func (service TenantService) GetAllWithLastName(lastName string) ([]*ent.Tenant, error) {
	return service.repo.GetAllWithLastName(lastName)
}

func (service TenantService) CreateTenant(jsonTenant []byte) (*ent.Tenant, error) {
	var tenant ent.Tenant

	if err := json.Unmarshal(jsonTenant, &tenant); err != nil {
		return nil, err
	}

	if err := tenant.Validate(); err != nil {
		return nil, err
	}

	dniAlreadyExists, err := service.repo.ExistsWithDNI(tenant.GetDNI())
	if err != nil {
		return nil, err
	}

	if dniAlreadyExists {
		return nil, ErrDuplicateDNI
	}

	emailAlreadyExists, err := service.repo.ExistsWithEmail(tenant.Email.String())
	if err != nil {
		return nil, err
	}

	if emailAlreadyExists {
		return nil, ErrDuplicateEmail
	}

	return service.repo.Save(&tenant)
}

func (service TenantService) ModifyByID(id int, requestBody []byte) (*ent.Tenant, error) {
	var newTenant ent.Tenant

	if err := json.Unmarshal(requestBody, &newTenant); err != nil {
		return nil, err
	}

	existingTenant, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}

	err = newTenant.DNI.Validate()
	if newTenant.DNI != 0 && err != nil {
		return nil, err
	}

	if newTenant.DNI != 0 {
		existingTenant.DNI = newTenant.DNI
	}

	err = newTenant.Email.Validate()
	if newTenant.Email != "" && err != nil {
		return nil, err
	}

	if newTenant.Email != "" {
		existingTenant.Email = newTenant.Email
	}

	return service.repo.Save(existingTenant)
}
