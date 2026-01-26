package integration

import (
	ent "cochera/domain/entities"
	"cochera/infra"
	"testing"
)

func TestPostgresSlotRepository_Save_Successfully_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	tenantsRepo := infra.NewPostgresTenantRepository(db)

	localTenant, _ := ent.NewTenant(1, 123, "Santi", "Ruiz", "Rosales 54", "+543333111111", "santiruiz@live.com", "12-2022")

	_, _ = tenantsRepo.Save(localTenant)

	slotsRepo := infra.NewPostgresSlotRepository(db)

	localSlot, _ := ent.NewSlot(1, 1, 1)

	savedSlot, err := slotsRepo.Save(localSlot)
	if err != nil {
		t.Fatal(err)
	}

	if *localSlot != *savedSlot {
		t.Fatal("Expected slot is different from saved slot")
	}
}

func TestPostgresSlotRepository_GetByID_Successfully_Integration(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	tenantsRepo := infra.NewPostgresTenantRepository(db)

	localTenant, _ := ent.NewTenant(1, 123, "Santi", "Ruiz", "Rosales 54", "+543333111111", "santiruiz@live.com", "12-2022")

	_, _ = tenantsRepo.Save(localTenant)

	slotsRepo := infra.NewPostgresSlotRepository(db)

	localSlot, _ := ent.NewSlot(6, 6, 1)

	_, _ = slotsRepo.Save(localSlot)

	savedSlot, err := slotsRepo.GetByID(6)
	if err != nil {
		t.Fatal(err)
	}

	if *localSlot != *savedSlot {
		t.Fatal("Expected slot is different from saved slot")
	}
}
