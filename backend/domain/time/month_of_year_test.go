package time

import "testing"

func TestMonthOfYearCreatedFromString(t *testing.T) {
	expectedMonthOfYear := "05-2001"
	monthOfYear, err := NewMonthOfYearFromString(expectedMonthOfYear)
	if err != nil {
		t.Error("Error instanciando un mes de año: ", err)
	}

	if monthOfYear.month != 5 {
		t.Errorf("Esperado mes 5, obtenido %v", monthOfYear.month)
	}

	if monthOfYear.year != 2001 {
		t.Errorf("Esperado año 2001, obtenido %v", monthOfYear.year)
	}
}
