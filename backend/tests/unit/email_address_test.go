package unit

import (
	"cochera/domain"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidEmailAddress(t *testing.T) {
	_, err := domain.NewEmailAddress("claude@speed.com")

	if err != nil {
		t.Fatal("Email no debe devolver error cuando es un string válido")
	}
}

func TestFailToCreateEmailAddressFromInvalidString(t *testing.T) {
	_, err := domain.NewEmailAddress("claudespeed.com")

	if err == nil || !errors.Is(err, domain.ErrInvalidEmailFormat) {
		t.Fatal("Email debe devolver error cuando no es un string válido")
	}
}

func TestFailToCreateEmailAddressFromStringTooLong(t *testing.T) {
	_, err := domain.NewEmailAddress("carl@johnson.commmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")

	if err == nil || !errors.Is(err, domain.ErrEmailTooLong) {
		t.Fatal("Email debe devolver error cuando es un string demasiado largo")
	}
}

func TestCreateValidEmptyEmailAddressFromNilValue(t *testing.T) {
	emailAddress, err := domain.NewEmailAddress(nil)

	if err != nil {
		t.Fatal("Email no debe devolver error cuando es un nil")
	}

	if emailAddress.Value != "" {
		t.Fatal("El valor del email debe ser vacío cuando es creado a partir de un nil")
	}
}

func TestFailToCreateEmailAddressFromNonStringType(t *testing.T) {
	_, err := domain.NewEmailAddress(map[string]int{"JL": 2579})

	if err == nil || !errors.Is(err, domain.ErrEmailMustBeString) {
		t.Fatal("Email debe devolver error cuando no es de tipo string")
	}
}

func TestCreateValidEmailAddressFromJSON(t *testing.T) {
	emailAddress := domain.EmailAddress{}
	err := emailAddress.UnmarshalJSON([]byte(`"huanglee@triads.org"`))

	if err != nil {
		t.Fatal("Email no debe devolver error cuando se decodifica un string válido desde json")
	}

	if emailAddress.Value != "huanglee@triads.org" {
		t.Fatal("Email decodificado de string desde json no coincide")
	}
}

func TestFailToCreateEmailAddressFromJSON(t *testing.T) {
	emailAddress := domain.EmailAddress{}
	err := emailAddress.UnmarshalJSON([]byte("777"))

	if err == nil || !errors.Is(err, domain.ErrEmailMustBeString) {
		t.Fatal("Email debe devolver error cuando se decodifica valor no string desde json")
	}

	if emailAddress.Value != "" {
		t.Fatal("El valor del Email no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeEmailAddressIntoJSON(t *testing.T) {
	emailAddress, _ := domain.NewEmailAddress("illo@juan.com")
	expectedJson := []byte(`"illo@juan.com"`)

	json, err := emailAddress.MarshalJSON()

	if err != nil {
		t.Fatal("No puede fallar la codificación del Email")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatal("Email codificado como json no coincide con esperado")
	}
}
