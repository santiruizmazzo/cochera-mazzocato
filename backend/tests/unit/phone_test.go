package unit

import (
	"cochera/domain"
	"errors"
	"reflect"
	"testing"
)

func TestCreateValidPhone(t *testing.T) {
	phone, err := domain.NewPhone("+543442407277")

	if err != nil {
		t.Fatal("Phone no debe devolver error cuando es un string válido")
	}

	if phone.CountryCode != "54" {
		t.Fatal("El código de país del teléfono no coincide con el esperado: ", phone.CountryCode)
	}

	if phone.LineNumber != "3442407277" {
		t.Fatal("El número de línea del teléfono no coincide con el esperado")
	}
}

func TestFailToCreatePhoneWithMissingCountryCode(t *testing.T) {
	_, err := domain.NewPhone("543442407277")

	if err == nil || !errors.Is(err, domain.ErrPhoneMustStartWithPlusSign) {
		t.Fatal("Phone debe devolver error cuando no tiene + al inicio")
	}
}

func TestCreateValidPhoneFromUruguayCodedNumber(t *testing.T) {
	phone, err := domain.NewPhone("+59812341234")

	if err != nil {
		t.Fatal("Phone no debe devolver error cuando es un string válido")
	}

	if phone.CountryCode != "598" {
		t.Fatal("El código de país del teléfono no coincide con el esperado: ", phone.CountryCode)
	}

	if phone.LineNumber != "12341234" {
		t.Fatal("El número de línea del teléfono no coincide con el esperado")
	}
}

func TestFailToCreatePhoneTooLong(t *testing.T) {
	_, err := domain.NewPhone("+5981234123412341234")

	if err == nil || !errors.Is(err, domain.ErrPhoneTooLong) {
		t.Fatal("Phone debe devolver error cuando es demasiado largo")
	}
}

func TestFailToCreatePhoneFullOfZeroes(t *testing.T) {
	_, err := domain.NewPhone("+0000000000")

	if err == nil || !errors.Is(err, domain.ErrPhoneFullOfZeroes) {
		t.Fatal("Phone debe devolver error cuando está lleno de ceros")
	}
}

func TestFailToCreatePhoneThatContainsLetters(t *testing.T) {
	_, err := domain.NewPhone("+57HolaComoEsta")

	if err == nil || !errors.Is(err, domain.ErrPhoneMustContainNumbersOnly) {
		t.Fatal("Phone debe devolver error cuando contiene caracteres")
	}
}

func TestFailToCreatePhoneThatIsNotAString(t *testing.T) {
	_, err := domain.NewPhone(482984292)

	if err == nil || !errors.Is(err, domain.ErrPhoneMustBeString) {
		t.Fatal("Phone debe devolver error cuando no es un string")
	}
}

func TestCreateEmptyPhoneFromNilValue(t *testing.T) {
	phone, err := domain.NewPhone(nil)

	if err != nil {
		t.Fatal("Phone no debe devolver error cuando se le pasa un nil")
	}

	if phone.CountryCode != "" || phone.LineNumber != "" {
		t.Fatal("El código de país un número deben estar vacíos cuando Phone se crea desde nil")
	}
}

func TestFailToCreatePhoneFromParaguayCodedNumber(t *testing.T) {
	_, err := domain.NewPhone("+595111111111")

	if err == nil || !errors.Is(err, domain.ErrPhoneUnsupportedCountryCode) {
		t.Fatal("Phone debe devolver error cuando no tiene un prefijo de país aceptado")
	}
}

func TestCreateValidPhoneFromJSON(t *testing.T) {
	phone := domain.Phone{}
	err := phone.UnmarshalJSON([]byte(`"+59844298511"`))

	if err != nil {
		t.Fatal("Phone no debe devolver error cuando se decodifica un string válido desde json")
	}

	if phone.CountryCode != "598" || phone.LineNumber != "44298511" {
		t.Fatal("Phone decodificado de string desde json no coincide")
	}
}

func TestFailToCreatePhoneFromJSON(t *testing.T) {
	phone := domain.Phone{}
	err := phone.UnmarshalJSON([]byte(`{"entry_month":"09-2025"}`))

	if err == nil || !errors.Is(err, domain.ErrPhoneMustBeString) {
		t.Fatal("Phone debe devolver error cuando se decodifica un tipo no string desde json")
	}

	if phone.CountryCode != "" || phone.LineNumber != "" {
		t.Fatal("El valor del Phone no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodePhoneIntoJSON(t *testing.T) {
	phone, _ := domain.NewPhone("+543442407276")
	expectedJson := []byte(`"+543442407276"`)

	json, err := phone.MarshalJSON()

	if err != nil {
		t.Fatal("No puede fallar la codificación del Phone")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatal("Phone codificado como json no coincide con esperado")
	}
}
