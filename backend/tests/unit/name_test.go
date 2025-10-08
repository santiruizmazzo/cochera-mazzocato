package unit

import (
	"cochera/domain"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidName(t *testing.T) {
	_, err := domain.NewName("Vaas")

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateNameFromNonStringType(t *testing.T) {
	_, err := domain.NewName(222)

	if err == nil || !errors.Is(err, domain.ErrNameMustBeString) {
		t.Fatalf("Name debe devolver error cuando no es un string")
	}
}

func TestFailToCreateNameTooLarge(t *testing.T) {
	_, err := domain.NewName("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	if err == nil || !errors.Is(err, domain.ErrNameTooLong) {
		t.Fatalf("Name debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidNameFromJSON(t *testing.T) {
	name := domain.Name("")
	err := name.UnmarshalJSON([]byte(`"Satoru"`))

	if err != nil {
		t.Fatalf("Name no debe devolver error cuando se decodifica un string válido desde json")
	}

	if name != "Satoru" {
		t.Fatalf("Name decodificado de string desde json no coincide")
	}
}

func TestFailToCreateNameFromJSON(t *testing.T) {
	name := domain.Name("")
	err := name.UnmarshalJSON([]byte(`{"id":2}`))

	if err == nil {
		t.Fatalf("Name debe devolver error cuando se decodifica valor no string desde json")
	}

	if name != "" {
		t.Fatalf("El valor del Name no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeNameIntoJSON(t *testing.T) {
	name, _ := domain.NewName("Big Smoke")
	expectedJson := []byte(`"Big Smoke"`)

	json, err := name.MarshalJSON()

	if err != nil {
		t.Fatalf("No puede fallar la codificación del Name")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatalf("Name codificado como json no coincide con esperado")
	}
}
