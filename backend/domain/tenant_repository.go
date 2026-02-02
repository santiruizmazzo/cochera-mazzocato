package domain

import (
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
)

type TenantRepository interface {
	GetByID(id int) (*ent.Tenant, error)
	GetAll() ([]*ent.Tenant, error)
	GetAllWithName(name string) ([]*ent.Tenant, error)
	GetAllWithLastName(lastName string) ([]*ent.Tenant, error)
	GetAllWithin(tenantIDs []vo.EntityID) ([]*ent.Tenant, error)
	Save(tenant *ent.Tenant) (*ent.Tenant, error)
	ExistsWithDNI(dni int) (bool, error)
	ExistsWithEmail(email string) (bool, error)
}
