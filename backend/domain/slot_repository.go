package domain

import ent "cochera/domain/entities"

type SlotRepository interface {
	Save(slot *ent.Slot) (*ent.Slot, error)
}
