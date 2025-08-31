package calendar

import "testing"

func TestMonthOfYearCreatedFromString(t *testing.T) {
	expectedMonthOfYear := "05-2001"
	monthOfYear, err := NewMonthOfYearFromString(expectedMonthOfYear)
	if err != nil {
		t.Fatal("Failed creating month of year: ", err)
	}

	if monthOfYear.month != 5 {
		t.Fatalf("Expected month 5, got %v", monthOfYear.month)
	}

	if monthOfYear.year != 2001 {
		t.Fatalf("Expected year 2001, got %v", monthOfYear.year)
	}
}
