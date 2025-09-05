package unit

import (
	"cochera/domain/calendar"
	"testing"
)

func TestMonthOfYearCreatedFromString(t *testing.T) {
	expectedMonthOfYear := "05-2001"
	monthOfYear, err := calendar.NewMonthOfYearFromString(expectedMonthOfYear)
	if err != nil {
		t.Fatal("Failed creating month of year: ", err)
	}

	if monthOfYear.Month != 5 {
		t.Fatalf("Expected month 5, got %v", monthOfYear.Month)
	}

	if monthOfYear.Year != 2001 {
		t.Fatalf("Expected year 2001, got %v", monthOfYear.Year)
	}
}
