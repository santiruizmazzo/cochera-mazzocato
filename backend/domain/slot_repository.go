package domain

import ent "cochera/domain/entities"

type SlotRepository interface {
	GetByID(id int) (*ent.Slot, error)
	GetAll() ([]*ent.Slot, error)
	Save(slot *ent.Slot) (*ent.Slot, error)
}
