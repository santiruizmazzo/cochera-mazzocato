package domain

import ent "cochera/domain/entities"

type TenantRepository interface {
	GetTenantByID(id int) (*ent.Tenant, error)
	GetAllTenants() ([]*ent.Tenant, error)
	GetAllTenantsByName(name string) ([]*ent.Tenant, error)
	GetAllTenantsByLastName(lastName string) ([]*ent.Tenant, error)
	Save(tenant *ent.Tenant) (*ent.Tenant, error)
	ExistsTenantWithDNI(dni uint32) (bool, error)
	ExistsTenantWithEmail(email string) (bool, error)
}
