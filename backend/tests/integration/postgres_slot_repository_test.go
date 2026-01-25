package integration

import (
	ent "cochera/domain/entities"
	"cochera/infra"
	"testing"
)

func TestPostgresSlotRepository_Save_Successfully_Integration(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip()
	}

	repo := infra.NewPostgresSlotRepository(db)

	localSlot, _ := ent.NewSlot(1, 1, 1)

	savedSlot, err := repo.Save(localSlot)
	if err != nil {
		t.Fatal(err)
	}

	if *localSlot != *savedSlot {
		t.Fatal("Expected slot is different from saved slot")
	}
}
