package services

import (
	"cochera/application/dtos"
	"cochera/domain"
	ent "cochera/domain/entities"
	vo "cochera/domain/value_objects"
)

type SlotService struct {
	repo domain.SlotRepository
}

func NewSlotService(repo domain.SlotRepository) *SlotService {
	return &SlotService{repo: repo}
}

func (service SlotService) UpdateSlot(id int, updateDTO dtos.UpdateSlotDTO) (*ent.Slot, error) {
	return &ent.Slot{ID: vo.EntityID(id), Number: *updateDTO.Number, TenantID: *updateDTO.TenantID}, nil
}
