package infra

import (
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
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

func (repo *InMemoryTenantRepository) GetAllWithin(tenantIDs []vo.EntityID) ([]*ent.Tenant, error) {
	panic("unimplemented")
}

func (repo InMemoryTenantRepository) GetByID(id int) (*ent.Tenant, error) {
	tenant := repo.Tenants[id]
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	return tenant, nil
}

func (repo InMemoryTenantRepository) Save(tenant *ent.Tenant) (*ent.Tenant, error) {
	if repo.err != nil {
		return nil, repo.err
	}
	tenant.ID = 1
	return tenant, nil
}

func (repo InMemoryTenantRepository) ExistsWithDNI(dni int) (bool, error) {
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
