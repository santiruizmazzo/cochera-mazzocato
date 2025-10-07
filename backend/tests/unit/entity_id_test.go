package unit

import (
	"cochera/domain"
	"errors"
	"testing"
)

func TestCreateValidEntityID(t *testing.T) {
	id, err := domain.NewEntityID(745)

	if err != nil {
		t.Fatal("Entity ID no debe devolver error cuando es un número válido")
	}

	if id != 745 {
		t.Fatal("Entity ID instanciado no coincide con esperado")
	}
}

func TestFailToCreateEntityIDFromNonNumericType(t *testing.T) {
	_, err := domain.NewEntityID("You shall not pass!")

	if err == nil || !errors.Is(err, domain.ErrIDMustBeANumber) {
		t.Fatal("Entity ID debe devolver error cuando no es un número")
	}
}
