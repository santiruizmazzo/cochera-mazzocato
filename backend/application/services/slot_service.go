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

func (service SlotService) GetByID(id int) (*ent.Slot, error) {
	return service.repo.GetByID(id)
}

func (service SlotService) UpdateSlot(id int, updateDTO dtos.UpdateSlotDTO) (*ent.Slot, error) {
	slot, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updateDTO.TenantID != nil {
		slot.SetTenantID(*updateDTO.TenantID)
	}

	return service.repo.Save(slot)
}
