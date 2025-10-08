package unit

import (
	"cochera/domain"
	"errors"
	"testing"
)

func TestMonthOfYearCreatedFromString(t *testing.T) {
	monthOfYear, err := domain.NewMonthOfYear("05-2001")

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
	_, err := domain.NewMonthOfYear("052001")

	if err == nil || !errors.Is(err, domain.ErrMonthOfYearInvalidFormat) {
		t.Fatal("Debería fallar la creación cuando el string no sigue el formato esperado")
	}
}

func TestFailToCreateMonthOfYearWithInvalidMonth(t *testing.T) {
	_, err := domain.NewMonthOfYear("20-2001")

	if err == nil || !errors.Is(err, domain.ErrMonthOfYearInvalidMonth) {
		t.Fatal("Debería fallar la creación cuando el mes no va del 1 al 12")
	}
}

func TestFailToCreateMonthOfYearWithInvalidYear(t *testing.T) {
	_, err := domain.NewMonthOfYear("01-20000")

	if err == nil || !errors.Is(err, domain.ErrMonthOfYearInvalidYear) {
		t.Fatal("Debería fallar la creación cuando el año no va del 0 al 9999")
	}
}

func TestFailToCreateMonthOfYearFromNonStringValue(t *testing.T) {
	_, err := domain.NewMonthOfYear(842892)

	if err == nil || !errors.Is(err, domain.ErrMonthOfYearMustBeString) {
		t.Fatal("Debería fallar la creación cuando no es un string")
	}
}
