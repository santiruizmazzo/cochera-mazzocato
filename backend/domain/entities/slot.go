package ent

import vo "cochera/domain/value_objects"

type Slot struct {
	ID       vo.EntityID   `json:"id"`
	Number   vo.SlotNumber `json:"number"`
	TenantID vo.EntityID   `json:"tenant_id"`
}
