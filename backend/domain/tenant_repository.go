package domain

import ent "cochera/domain/entities"

type TenantRepository interface {
	GetByID(id int) (*ent.Tenant, error)
	GetAll() ([]*ent.Tenant, error)
	GetAllWithName(name string) ([]*ent.Tenant, error)
	GetAllWithLastName(lastName string) ([]*ent.Tenant, error)
	Save(tenant *ent.Tenant) (*ent.Tenant, error)
	ExistsTenantWithDNI(dni uint32) (bool, error)
	ExistsTenantWithEmail(email string) (bool, error)
}
