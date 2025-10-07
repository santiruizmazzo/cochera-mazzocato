package unit

import (
	"cochera/domain"
	"errors"
	"testing"
)

func TestCreateValidPhone(t *testing.T) {
	phone, err := domain.NewPhone("+543442407277")

	if err != nil {
		t.Fatalf("Phone no debe devolver error cuando es un string válido")
	}

	if phone.CountryCode != "54" {
		t.Fatal("El código de país del teléfono no coincide con el esperado: ", phone.CountryCode)
	}

	if phone.LineNumber != "3442407277" {
		t.Fatalf("El número de línea del teléfono no coincide con el esperado")
	}
}

func TestFailToCreatePhoneWithMissingCountryCode(t *testing.T) {
	_, err := domain.NewPhone("543442407277")

	if err == nil || !errors.Is(err, domain.ErrPhoneMustStartWithPlusSign) {
		t.Fatalf("Phone debe devolver error cuando no tiene + al inicio")
	}
}

func TestCreateValidPhoneFromUruguayCodedNumber(t *testing.T) {
	phone, err := domain.NewPhone("+59812341234")

	if err != nil {
		t.Fatalf("Phone no debe devolver error cuando es un string válido")
	}

	if phone.CountryCode != "598" {
		t.Fatal("El código de país del teléfono no coincide con el esperado: ", phone.CountryCode)
	}

	if phone.LineNumber != "12341234" {
		t.Fatalf("El número de línea del teléfono no coincide con el esperado")
	}
}

func TestFailToCreatePhoneTooLong(t *testing.T) {
	_, err := domain.NewPhone("+5981234123412341234")

	if err == nil || !errors.Is(err, domain.ErrPhoneTooLong) {
		t.Fatalf("Phone debe devolver error cuando es demasiado largo")
	}
}

func TestFailToCreatePhoneFullOfZeroes(t *testing.T) {
	_, err := domain.NewPhone("+0000000000")

	if err == nil || !errors.Is(err, domain.ErrPhoneFullOfZeroes) {
		t.Fatalf("Phone debe devolver error cuando está lleno de ceros")
	}
}

func TestFailToCreatePhoneThatContainsLetters(t *testing.T) {
	_, err := domain.NewPhone("+57HolaComoEsta")

	if err == nil || !errors.Is(err, domain.ErrPhoneMustContainNumbersOnly) {
		t.Fatalf("Phone debe devolver error cuando contiene caracteres")
	}
}
