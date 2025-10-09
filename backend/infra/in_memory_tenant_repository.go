package infra

import (
	"cochera/domain"
)

type InMemoryTenantRepository struct {
	Tenants map[int]*domain.Tenant
	err     error
}

func (repo InMemoryTenantRepository) GetAllTenants() ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range repo.Tenants {
		list = append(list, tenant)
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetAllTenantsByName(name string) ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range repo.Tenants {
		if tenant.HasName(name) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetAllTenantsByLastName(lastName string) ([]*domain.Tenant, error) {
	var list []*domain.Tenant
	for _, tenant := range repo.Tenants {
		if tenant.HasLastName(lastName) {
			list = append(list, tenant)
		}
	}
	return list, nil
}

func (repo InMemoryTenantRepository) GetTenantByID(id int) (*domain.Tenant, error) {
	return nil, ErrTenantNotFound
}

func (repo InMemoryTenantRepository) Save(tenant *domain.Tenant) (*domain.Tenant, error) {
	if repo.err != nil {
		return nil, repo.err
	}
	tenant.ID = 1
	return tenant, nil
}

func (repo InMemoryTenantRepository) ExistsTenantWithDNI(dni uint32) (bool, error) {
	for _, tenant := range repo.Tenants {
		if tenant != nil && tenant.HasDNI(dni) {
			return true, nil
		}
	}
	return false, nil
}

func (repo InMemoryTenantRepository) ExistsTenantWithEmail(email string) (bool, error) {
	for _, tenant := range repo.Tenants {
		if tenant != nil && tenant.HasEmail(email) {
			return true, nil
		}
	}
	return false, nil
}
