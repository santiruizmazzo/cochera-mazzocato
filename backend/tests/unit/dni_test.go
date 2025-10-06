package unit

import (
	"cochera/domain"
	"testing"
)

func TestCreateValidDNI(t *testing.T) {
	_, err := domain.NewDNI(12345678)

	if err != nil {
		t.Fatalf("DNI no debe devolver error cuando es un número válido")
	}
}
