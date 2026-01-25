package services

import (
	"cochera/application/dtos"
	"cochera/domain"
	ent "cochera/domain/entities"
)

type SlotService struct {
	repo domain.SlotRepository
}

func NewSlotService(repo domain.SlotRepository) *SlotService {
	return &SlotService{repo: repo}
}

func (service SlotService) UpdateSlot(id int, updateDTO dtos.UpdateSlotDTO) (*ent.Slot, error) {
	return nil, nil
}
