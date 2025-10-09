package unit

import (
	vo "cochera/domain/value_objects"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidName(t *testing.T) {
	_, err := vo.NewName("Vaas")

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateNameFromNonStringType(t *testing.T) {
	_, err := vo.NewName(222)

	if err == nil || !errors.Is(err, vo.ErrNameMustBeString) {
		t.Fatalf("Name debe devolver error cuando no es un string")
	}
}

func TestFailToCreateNameTooLarge(t *testing.T) {
	_, err := vo.NewName("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	if err == nil || !errors.Is(err, vo.ErrNameTooLong) {
		t.Fatalf("Name debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidNameFromJSON(t *testing.T) {
	name := vo.Name("")
	err := name.UnmarshalJSON([]byte(`"Satoru"`))

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando se decodifica un string válido desde json")
	}

	if name != "Satoru" {
		t.Fatalf("Name decodificado de string desde json no coincide")
	}
}

func TestFailToCreateNameFromJSON(t *testing.T) {
	name := vo.Name("")
	err := name.UnmarshalJSON([]byte(`{"id":2}`))

	if err == nil {
		t.Fatalf("Name debe devolver error cuando se decodifica valor no string desde json")
	}

	if name != "" {
		t.Fatalf("El valor del Name no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeNameIntoJSON(t *testing.T) {
	name, _ := vo.NewName("Big Smoke")
	expectedJson := []byte(`"Big Smoke"`)

	json, err := name.MarshalJSON()

	if err != nil {
		t.Fatalf("No puede fallar la codificación del Name")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatalf("Name codificado como json no coincide con esperado")
	}
}

func TestFailToCreateNameFromStringWithNumbers(t *testing.T) {
	_, err := vo.NewName("askflj98242")

	if err == nil || !errors.Is(err, vo.ErrNameMustIncludeCharactersOnly) {
		t.Fatalf("Name debe devolver error cuando contiene números")
	}
}
