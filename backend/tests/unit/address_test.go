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

func TestFailToCreateAddressFromStringTooLong(t *testing.T) {
	_, err := domain.NewAddress("KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK")

	if err == nil {
		t.Fatalf("Address debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidAddressFromJSON(t *testing.T) {
	address := domain.Address{}
	err := address.UnmarshalJSON([]byte(`"742 Evergreen Terrace, Springfield"`))

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando se decodifica un string válido desde json")
	}

	if address.Value != "742 Evergreen Terrace, Springfield" {
		t.Fatalf("Address decodificado de string desde json no coincide")
	}
}

func TestFailToCreateAddressFromJSON(t *testing.T) {
	address := domain.Address{}
	err := address.UnmarshalJSON([]byte("130501"))

	if err == nil {
		t.Fatalf("Address debe devolver error cuando se decodifica valor no string desde json")
	}

	if address.Value != "" {
		t.Fatalf("El valor del Address no decodificado correctamente debe ser vacío")
	}
}
