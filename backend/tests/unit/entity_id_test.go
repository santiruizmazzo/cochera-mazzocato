package unit

import (
	"cochera/domain"
	"errors"
	"math"
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

	if err == nil || !errors.Is(err, domain.ErrIDMustBeAnInteger) {
		t.Fatal("Entity ID debe devolver error cuando no es un número entero")
	}
}

func TestFailToCreateEntityIDFromTooSmallValue(t *testing.T) {
	_, err := domain.NewEntityID(0)

	if err == nil || !errors.Is(err, domain.ErrIDTooSmall) {
		t.Fatal("Entity ID debe devolver error cuando es menor a 1")
	}
}

func TestFailToCreateEntityIDFromTooBigValue(t *testing.T) {
	_, err := domain.NewEntityID(math.MaxUint32 + 10)

	if err == nil || !errors.Is(err, domain.ErrIDTooBig) {
		t.Fatal("Entity ID debe devolver error cuando mayor al limite del tipo de dato")
	}
}
