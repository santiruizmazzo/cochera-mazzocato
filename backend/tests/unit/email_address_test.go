package unit

import (
	"cochera/domain"
	"errors"
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
