package unit

import (
	"cochera/domain"
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
