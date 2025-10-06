package unit

import (
	"cochera/domain"
	"errors"
	"testing"
)

func TestCreateValidDNI(t *testing.T) {
	_, err := domain.NewDNI(12345678)

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando es un número válido")
	}
}

func TestCreateValidDNIFromFloatNumber(t *testing.T) {
	dni, err := domain.NewDNI(12345678.8298)

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando es un número válido")
	}

	if dni.Value != 12345678 {
		t.Fatalf("El valor del DNI debería igual a la parte entera del float")
	}
}

func TestFailToCreateDNIFromZero(t *testing.T) {
	_, err := domain.NewDNI(0)

	if err == nil || !errors.Is(err, domain.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es menor a 1")
	}
}

func TestFailToCreateDNIFromOverflowValue(t *testing.T) {
	_, err := domain.NewDNI(4500000000)

	if err == nil || !errors.Is(err, domain.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es mayor a 4294967295")
	}
}

func TestFailToCreateDNIFromOverflowFloatValue(t *testing.T) {
	_, err := domain.NewDNI(4500000000.01)

	if err == nil || !errors.Is(err, domain.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es mayor a 4294967295")
	}
}

func TestFailToCreateDNIFromNegativeValue(t *testing.T) {
	_, err := domain.NewDNI(-99)

	if err == nil || !errors.Is(err, domain.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es menor a 1")
	}
}

func TestFailToCreateDNIFromNegativeFloatValue(t *testing.T) {
	_, err := domain.NewDNI(-99.849)

	if err == nil || !errors.Is(err, domain.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es menor a 1")
	}
}
