package unit

import (
	vo "cochera/domain/value_objects"
	"reflect"
	"testing"
)

func TestCreateValidAddress(t *testing.T) {
	_, err := vo.NewAddress("Roseti 745, Chacarita, CABA")

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando es un string válido")
	}
}

func TestCreateValidAddressFromNilValue(t *testing.T) {
	_, err := vo.NewAddress(nil)

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateAddressFromNonStringType(t *testing.T) {
	_, err := vo.NewAddress([]byte("blabla"))

	if err == nil {
		t.Fatalf("Address debe devolver error cuando no es un string")
	}
}

func TestFailToCreateAddressFromStringTooLong(t *testing.T) {
	_, err := vo.NewAddress("KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK")

	if err == nil {
		t.Fatalf("Address debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidAddressFromJSON(t *testing.T) {
	address := vo.Address("")
	err := address.UnmarshalJSON([]byte(`"742 Evergreen Terrace, Springfield"`))

	if err != nil {
		t.Fatalf("Address no debe devolver error cuando se decodifica un string válido desde json")
	}

	if address != "742 Evergreen Terrace, Springfield" {
		t.Fatalf("Address decodificado de string desde json no coincide")
	}
}

func TestFailToCreateAddressFromJSON(t *testing.T) {
	address := vo.Address("")
	err := address.UnmarshalJSON([]byte("130501"))

	if err == nil {
		t.Fatalf("Address debe devolver error cuando se decodifica valor no string desde json")
	}

	if address != "" {
		t.Fatalf("El valor del Address no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeAddressIntoJSON(t *testing.T) {
	address, _ := vo.NewAddress("Av. Paseo Colón 850")
	expectedJson := []byte(`"Av. Paseo Colón 850"`)

	json, err := address.MarshalJSON()

	if err != nil {
		t.Fatalf("No puede fallar la codificación del Address")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatalf("Address codificado como json no coincide con esperado")
	}
}
