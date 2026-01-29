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

func (service SlotService) GetAll() ([]*ent.Slot, error) {
	return nil, nil
}

func (service SlotService) UpdateSlot(id int, updateDTO dtos.UpdateSlotDTO) (*ent.Slot, error) {
	slot, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}

	var tenantID int
	if updateDTO.TenantID != nil {
		tenantID = int(*updateDTO.TenantID)

		_, err = service.tenantRepo.GetByID(tenantID)
		if err != nil {
			return nil, err
		}

	}
	slot.SetTenantID(tenantID)

	return service.slotRepo.Save(slot)
}
