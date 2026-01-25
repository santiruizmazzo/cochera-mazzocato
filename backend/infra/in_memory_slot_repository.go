package infra

import ent "cochera/domain/entities"

type InMemorySlotRepository struct {
	Slots map[int]*ent.Slot
	// err   error
}
