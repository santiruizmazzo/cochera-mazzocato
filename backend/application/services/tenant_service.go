package services

import (
	"cochera/application/dtos"
	"cochera/domain"
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
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

func (service TenantService) GetAllWithin(tenantIDs []vo.EntityID) ([]*ent.Tenant, error) {
	return service.repo.GetAllWithin(tenantIDs)
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

func (service TenantService) UpdateTenant(id int, updateDTO dtos.UpdateTenantDTO) (*ent.Tenant, error) {
	tenant, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updateDTO.DNI != nil && *updateDTO.DNI != tenant.DNI {
		dniAlreadyExists, err := service.repo.ExistsWithDNI(int(*updateDTO.DNI))
		if err != nil {
			return nil, err
		}
		if dniAlreadyExists {
			return nil, ErrDuplicateDNI
		}
		tenant.SetDNI(*updateDTO.DNI)
	}

	if updateDTO.Name != nil {
		tenant.SetName(*updateDTO.Name)
	}

	if updateDTO.LastName != nil {
		tenant.SetLastName(*updateDTO.LastName)
	}

	if updateDTO.Address != nil {
		tenant.SetAddress(*updateDTO.Address)
	}

	if updateDTO.Phone != nil {
		tenant.SetPhone(*updateDTO.Phone)
	}

	if updateDTO.Email != nil && *updateDTO.Email != tenant.Email {
		emailAlreadyUsed, err := service.repo.ExistsWithEmail(string(*updateDTO.Email))
		if err != nil {
			return nil, err
		}
		if emailAlreadyUsed {
			return nil, ErrDuplicateEmail
		}
		tenant.SetEmail(*updateDTO.Email)
	}

	if updateDTO.EntryMonth != nil {
		tenant.SetEntryMonth(*updateDTO.EntryMonth)
	}

	return service.repo.Save(tenant)
}
