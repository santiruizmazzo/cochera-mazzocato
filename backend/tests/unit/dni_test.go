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
		t.Fatalf("El valor del DNI debería ser igual a la parte entera del float")
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

func TestFailToCreateDNIFromNonNumericValue(t *testing.T) {
	_, err := domain.NewDNI("hola")

	if err == nil || !errors.Is(err, domain.ErrDNIMustBeANumber) {
		t.Fatalf("DNI debe devolver error cuando no es un número")
	}
}

func TestCreateValidDNIFromJSON(t *testing.T) {
	dni := domain.DNI{}
	err := dni.UnmarshalJSON([]byte("43295798"))

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando se decodifica un número válido desde json")
	}

	if dni.Value != 43295798 {
		t.Fatalf("DNI decodificado de entero desde json no coincide")
	}
}

func TestFailToCreateDNIFromJSON(t *testing.T) {
	dni := domain.DNI{}
	err := dni.UnmarshalJSON([]byte(`{"name":"Roberto"}`))

	if err == nil {
		t.Fatalf("DNI debe devolver error cuando se decodifica valor no numérico desde json")
	}

	if dni.Value != 0 {
		t.Fatalf("El valor del DNI no decodificado correctamente debe ser cero")
	}
}
