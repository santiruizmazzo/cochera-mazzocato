package unit

import (
	vo "cochera/domain/value_objects"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidDNI(t *testing.T) {
	_, err := vo.NewDNI(12345678)

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando es un número válido")
	}
}

func TestFailToCreateValidDNIFromFloatNumber(t *testing.T) {
	_, err := vo.NewDNI(12345678.8298)

	if err == nil || !errors.Is(err, vo.ErrDNIMustBeAnInteger) {
		t.Fatal("DNI debe devolver error cuando no es un número entero")
	}
}

func TestFailToCreateDNIFromZero(t *testing.T) {
	_, err := vo.NewDNI(0)

	if err == nil || !errors.Is(err, vo.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es menor a 1")
	}
}

func TestFailToCreateDNIFromOverflowValue(t *testing.T) {
	_, err := vo.NewDNI(4500000000)

	if err == nil || !errors.Is(err, vo.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es mayor a 4294967295")
	}
}

func TestFailToCreateDNIFromNegativeValue(t *testing.T) {
	_, err := vo.NewDNI(-99)

	if err == nil || !errors.Is(err, vo.ErrDNINotInValidRange) {
		t.Fatalf("DNI debe devolver error cuando es menor a 1")
	}
}

func TestFailToCreateDNIFromNonNumericValue(t *testing.T) {
	_, err := vo.NewDNI("hola")

	if err == nil || !errors.Is(err, vo.ErrDNIMustBeAnInteger) {
		t.Fatalf("DNI debe devolver error cuando no es un número")
	}
}

func TestCreateValidDNIFromJSON(t *testing.T) {
	dni := vo.DNI(0)
	err := dni.UnmarshalJSON([]byte("43295798"))

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando se decodifica un número válido desde json")
	}

	if dni != 43295798 {
		t.Fatalf("DNI decodificado de entero desde json no coincide")
	}
}

func TestFailToCreateDNIFromJSON(t *testing.T) {
	dni := vo.DNI(0)
	err := dni.UnmarshalJSON([]byte(`{"name":"Roberto"}`))

	if err == nil {
		t.Fatalf("DNI debe devolver error cuando se decodifica valor no numérico desde json")
	}

	if dni != 0 {
		t.Fatalf("El valor del DNI no decodificado correctamente debe ser cero")
	}
}

func TestSuccessfullyEncodeDNIIntoJSON(t *testing.T) {
	dni, _ := vo.NewDNI(666)
	expectedJson := []byte("666")

	json, err := dni.MarshalJSON()

	if err != nil {
		t.Fatalf("No puede fallar la codificación del DNI")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatalf("DNI codificado como json no coincide con esperado")
	}
}
