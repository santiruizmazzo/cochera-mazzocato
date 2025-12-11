package infra

import (
	ent "cochera/domain/entities"
	"encoding/json"
)

type InMemoryTenantRepository struct {
	Tenants map[int]*ent.Tenant
	err     error
}

func (repo InMemoryTenantRepository) GetAll() ([]*ent.Tenant, error) {
	var list []*ent.Tenant
	for _, tenant := range repo.Tenants {
		list = append(list, tenant)
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetAllWithName(name string) ([]*ent.Tenant, error) {
	var list []*ent.Tenant
	for _, tenant := range repo.Tenants {
		if tenant.HasName(name) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetAllWithLastName(lastName string) ([]*ent.Tenant, error) {
	var list []*ent.Tenant
	for _, tenant := range repo.Tenants {
		if tenant.HasLastName(lastName) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetByID(id int) (*ent.Tenant, error) {
	return nil, ErrTenantNotFound
}

func (repo InMemoryTenantRepository) Save(tenant *ent.Tenant) (*ent.Tenant, error) {
	if repo.err != nil {
		return nil, repo.err
	}
	tenant.ID = 1
	return tenant, nil
}

func (repo InMemoryTenantRepository) ExistsWithDNI(dni uint32) (bool, error) {
	for _, tenant := range repo.Tenants {
		if tenant != nil && tenant.HasDNI(dni) {
			return true, nil
		}
	}
	return false, nil
}

func (repo InMemoryTenantRepository) ExistsWithEmail(email string) (bool, error) {
	for _, tenant := range repo.Tenants {
		if tenant != nil && tenant.HasEmail(email) {
			return true, nil
		}
	}
	return false, nil
}

func (repo InMemoryTenantRepository) ModifyByID(id int, requestBody []byte) (*ent.Tenant, error) {
	var tenant ent.Tenant

	if err := json.Unmarshal(requestBody, &tenant); err != nil {
		return nil, err
	}

	existingTenant := repo.Tenants[id]
	existingTenant.DNI = tenant.DNI

	return existingTenant, nil
}
