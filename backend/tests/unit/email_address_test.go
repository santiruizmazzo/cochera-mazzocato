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
