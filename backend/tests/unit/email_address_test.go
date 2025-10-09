package unit

import (
	vo "cochera/domain/value_objects"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidEmailAddress(t *testing.T) {
	_, err := vo.NewEmailAddress("claude@speed.com")

	if err != nil {
		t.Fatal("Email no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateEmailAddressFromInvalidString(t *testing.T) {
	_, err := vo.NewEmailAddress("claudespeed.com")

	if err == nil || !errors.Is(err, vo.ErrInvalidEmailFormat) {
		t.Fatal("Email debe devolver error cuando no es un string válido")
	}
}

func TestFailToCreateEmailAddressFromStringTooLong(t *testing.T) {
	_, err := vo.NewEmailAddress("carl@johnson.commmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")

	if err == nil || !errors.Is(err, vo.ErrEmailTooLong) {
		t.Fatal("Email debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidEmptyEmailAddressFromNilValue(t *testing.T) {
	emailAddress, err := vo.NewEmailAddress(nil)

	if err != nil {
		t.Fatal("Email no debe devolver error cuando es un nil")
	}

	if emailAddress != "" {
		t.Fatal("El valor del email debe ser vacío cuando es creado a partir de un nil")
	}
}

func TestFailToCreateEmailAddressFromNonStringType(t *testing.T) {
	_, err := vo.NewEmailAddress(map[string]int{"JL": 2579})

	if err == nil || !errors.Is(err, vo.ErrEmailMustBeString) {
		t.Fatal("Email debe devolver error cuando no es de tipo string")
	}
}

func TestCreateValidEmailAddressFromJSON(t *testing.T) {
	emailAddress := vo.EmailAddress("")
	err := emailAddress.UnmarshalJSON([]byte(`"huanglee@triads.org"`))

	if err != nil {
		t.Fatal("Email no debe devolver error cuando se decodifica un string válido desde json")
	}

	if emailAddress != "huanglee@triads.org" {
		t.Fatal("Email decodificado de string desde json no coincide")
	}
}

func TestFailToCreateEmailAddressFromJSON(t *testing.T) {
	emailAddress := vo.EmailAddress("")
	err := emailAddress.UnmarshalJSON([]byte("777"))

	if err == nil || !errors.Is(err, vo.ErrEmailMustBeString) {
		t.Fatal("Email debe devolver error cuando se decodifica valor no string desde json")
	}

	if emailAddress != "" {
		t.Fatal("El valor del Email no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeEmailAddressIntoJSON(t *testing.T) {
	emailAddress, _ := vo.NewEmailAddress("illo@juan.com")
	expectedJson := []byte(`"illo@juan.com"`)

	json, err := emailAddress.MarshalJSON()

	if err != nil {
		t.Fatal("No puede fallar la codificación del Email")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatal("Email codificado como json no coincide con esperado")
	}
}
