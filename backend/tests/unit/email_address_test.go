package unit

import (
	"cochera/domain"
	"testing"
)

func TestCreateValidEmailAddress(t *testing.T) {
	_, err := domain.NewEmailAddress("claude@speed.com")

	if err != nil {
		t.Fatalf("Email no debe devolver error cuando es un string válido")
	}
}
