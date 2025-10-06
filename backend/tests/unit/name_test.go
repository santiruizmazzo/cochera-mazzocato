package unit

import (
	"cochera/domain"
	"testing"
)

func TestCreateValidName(t *testing.T) {
	_, err := domain.NewName("Vaas")

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando es un string válido")
	}
}
