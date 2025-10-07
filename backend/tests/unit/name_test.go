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

func TestFailToCreateNameTooLarge(t *testing.T) {
	_, err := domain.NewName("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	if err == nil || !errors.Is(err, domain.ErrNameTooLong) {
		t.Fatalf("Name debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidNameFromJSON(t *testing.T) {
	name := domain.Name{}
	err := name.UnmarshalJSON([]byte(`"Satoru"`))

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando se decodifica un string válido desde json")
	}

	if name.Value != "Satoru" {
		t.Fatalf("Name decodificado de string desde json no coincide")
	}
}

func TestFailToCreateNameFromJSON(t *testing.T) {
	name := domain.Name{}
	err := name.UnmarshalJSON([]byte(`{"id":2}`))

	if err == nil {
		t.Fatalf("Name debe devolver error cuando se decodifica valor no string desde json")
	}

	if name.Value != "" {
		t.Fatalf("El valor del Name no decodificado correctamente debe ser vacío")
	}
}
