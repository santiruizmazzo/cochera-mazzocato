package unit

import (
	"cochera/domain"
	"testing"
)

func TestCreateValidAddress(t *testing.T) {
	_, err := domain.NewAddress("Roseti 745, Chacarita, CABA")

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando es un string válido")
	}
}

func TestCreateValidAddressFromNilValue(t *testing.T) {
	_, err := domain.NewAddress(nil)

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateAddressFromNonStringType(t *testing.T) {
	_, err := domain.NewAddress([]byte("blabla"))

	if err == nil {
		t.Fatalf("Address debe devolver error cuando no es un string")
	}
}
