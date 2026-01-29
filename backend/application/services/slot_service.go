package services

import (
	"cochera/application/dtos"
	"cochera/domain"
	ent "cochera/domain/entities"
)

type SlotService struct {
	slotRepo   domain.SlotRepository
	tenantRepo domain.TenantRepository
}

func NewSlotService(slotRepo domain.SlotRepository, tenantRepo domain.TenantRepository) *SlotService {
	return &SlotService{slotRepo: slotRepo, tenantRepo: tenantRepo}
}

func (service SlotService) GetByID(id int) (*ent.Slot, error) {
	return service.slotRepo.GetByID(id)
}

func (service SlotService) UpdateSlot(id int, updateDTO dtos.UpdateSlotDTO) (*ent.Slot, error) {
	slot, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updateDTO.TenantID != nil {
		_, err = service.tenantRepo.GetByID(int(*updateDTO.TenantID))
		if err != nil {
			return nil, err
		}

		slot.SetTenantID(*updateDTO.TenantID)
	}

	return service.slotRepo.Save(slot)
}
