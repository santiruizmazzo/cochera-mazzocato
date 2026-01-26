package ent

import (
	vo "cochera/domain/value_objects"
)

type Slot struct {
	ID       vo.EntityID   `json:"id"`
	Number   vo.SlotNumber `json:"number"`
	TenantID vo.EntityID   `json:"tenant_id"`
}

func NewSlot(id int, number int, tenantID int) (*Slot, error) {
	return &Slot{ID: vo.EntityID(id), Number: vo.SlotNumber(number), TenantID: vo.EntityID(tenantID)}, nil
}
