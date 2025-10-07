package unit

import (
	"cochera/domain"
	"testing"
)

func TestCreateValidEntityID(t *testing.T) {
	id, err := domain.NewEntityID(745)

	if err != nil {
		t.Fatal("Entity ID no debe devolver error cuando es un número válido")
	}

	if id != 745 {
		t.Fatal("Entity ID instanciado no coincide con esperado")
	}
}
