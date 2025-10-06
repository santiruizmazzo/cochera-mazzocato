package unit

import (
	"cochera/domain"
	"errors"
	"testing"
)

func TestCreateValidName(t *testing.T) {
	_, err := domain.NewName("Vaas")

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateNameFromNonStringType(t *testing.T) {
	_, err := domain.NewName(222)

	if err == nil || !errors.Is(err, domain.ErrNameMustBeAString) {
		t.Fatalf("Name debe devolver error cuando no es un string")
	}
}
