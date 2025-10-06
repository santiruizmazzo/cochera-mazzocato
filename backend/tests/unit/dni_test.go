package unit

import (
	"cochera/domain"
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
