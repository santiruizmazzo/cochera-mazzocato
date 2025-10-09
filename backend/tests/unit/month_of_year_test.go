package unit

import (
	vo "cochera/domain/value_objects"
	"errors"
	"reflect"
	"testing"
)

func TestMonthOfYearCreatedFromString(t *testing.T) {
	monthOfYear, err := vo.NewMonthOfYear("05-2001")

	if err != nil {
		t.Fatal("Error al crear MonthOfYear: ", err)
	}

	if monthOfYear.Month != 5 {
		t.Fatal("Esperado mes 5, obtenido ", monthOfYear.Month)
	}

	if monthOfYear.Year != 2001 {
		t.Fatal("Esperado año 2001, obtenido ", monthOfYear.Year)
	}
}

func TestFailToCreateMonthOfYearFromInvalidFormatString(t *testing.T) {
	_, err := vo.NewMonthOfYear("052001")

	if err == nil || !errors.Is(err, vo.ErrMonthOfYearInvalidFormat) {
		t.Fatal("Debería fallar la creación cuando el string no sigue el formato esperado")
	}
}

func TestFailToCreateMonthOfYearWithInvalidMonth(t *testing.T) {
	_, err := vo.NewMonthOfYear("20-2001")

	if err == nil || !errors.Is(err, vo.ErrMonthOfYearInvalidMonth) {
		t.Fatal("Debería fallar la creación cuando el mes no va del 1 al 12")
	}
}

func TestFailToCreateMonthOfYearWithInvalidYear(t *testing.T) {
	_, err := vo.NewMonthOfYear("01-20000")

	if err == nil || !errors.Is(err, vo.ErrMonthOfYearInvalidYear) {
		t.Fatal("Debería fallar la creación cuando el año no va del 0 al 9999")
	}
}

func TestFailToCreateMonthOfYearFromNonStringValue(t *testing.T) {
	_, err := vo.NewMonthOfYear(842892)

	if err == nil || !errors.Is(err, vo.ErrMonthOfYearMustBeString) {
		t.Fatal("Debería fallar la creación cuando no es un string")
	}
}

func TestCreateValidMonthOfYearFromJSON(t *testing.T) {
	monthOfYear := vo.MonthOfYear{}
	err := monthOfYear.UnmarshalJSON([]byte(`"09-2001"`))

	if err != nil {
		t.Fatal("MonthOfYear no debe devolver error cuando se decodifica un string válido desde json")
	}

	if monthOfYear.String() != "09-2001" {
		t.Fatal("MonthOfYear decodificado de string desde json no coincide")
	}
}

func TestFailToCreateMonthOfYearFromJSON(t *testing.T) {
	monthOfYear := vo.MonthOfYear{}
	err := monthOfYear.UnmarshalJSON([]byte("[1,2,3]"))

	if err == nil {
		t.Fatal("MonthOfYear debe devolver error cuando se decodifica valor no string desde json")
	}

	if monthOfYear.Month != 0 && monthOfYear.Year != 0 {
		t.Fatal("El valor del MonthOfYear no decodificado correctamente debe ser vacío")
	}
}

func TestSuccessfullyEncodeMonthOfYearIntoJSON(t *testing.T) {
	monthOfYear, _ := vo.NewMonthOfYear("04-1999")
	expectedJson := []byte(`"04-1999"`)

	json, err := monthOfYear.MarshalJSON()

	if err != nil {
		t.Fatal("No puede fallar la codificación del MonthOfYear")
	}

	if !reflect.DeepEqual(json, expectedJson) {
		t.Fatal("MonthOfYear codificado como json no coincide con esperado")
	}
}
