package infra

import ent "cochera/domain/entities"

type InMemorySlotRepository struct {
	Slots map[int]*ent.Slot
	err   error
}

func (repo InMemorySlotRepository) GetByID(id int) (*ent.Slot, error) {
	slot := repo.Slots[id]
	if slot == nil {
		return nil, ErrSlotNotFound
	}

	return slot, nil
}

func (repo InMemorySlotRepository) Save(slot *ent.Slot) (*ent.Slot, error) {
	if repo.err != nil {
		return nil, repo.err
	}
	return slot, nil
}
