package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MonthOfYear struct {
	Month uint8
	Year  uint16
}

var (
	ErrMonthOfYearInvalidFormat = errors.New(`un mes de año debe seguir el formato: "MM-YYYY"`)
	ErrMonthOfYearMustBeString  = errors.New(`un mes de año debe ser un string con formato: "MM-YYYY"`)
	ErrMonthNotParseable        = errors.New("imposible procesar este mes")
	ErrYearNotParseable         = errors.New("imposible procesar este año")
)

func NewMonthOfYear(rawValue any) (MonthOfYear, error) {
	var stringValue string
	var ok bool

	if stringValue, ok = rawValue.(string); !ok {
		return MonthOfYear{}, ErrMonthOfYearMustBeString
	}

	parts := strings.Split(stringValue, "-")
	if len(parts) != 2 {
		return MonthOfYear{}, ErrMonthOfYearInvalidFormat
	}

	month, err := ParseMonth(parts[0])
	if err != nil {
		return MonthOfYear{}, err
	}

	year, err := ParseYear(parts[1])

	return MonthOfYear{Month: month, Year: year}, err
}

func ParseMonth(s string) (uint8, error) {
	newInt, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		err = ErrMonthNotParseable
	}

	return uint8(newInt), err
}

func ParseYear(s string) (uint16, error) {
	newInt, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		err = ErrYearNotParseable
	}

	return uint16(newInt), err
}

func (monthOfYear MonthOfYear) String() string {
	return fmt.Sprintf("%02d-%04d", monthOfYear.Month, monthOfYear.Year)
}

func (monthOfYear *MonthOfYear) UnmarshalJSON(data []byte) error {
	var stringFormatMonth string
	if err := json.Unmarshal(data, &stringFormatMonth); err != nil {
		return err
	}

	newMonthOfYear, err := NewMonthOfYear(stringFormatMonth)
	if err != nil {
		return err
	}

	*monthOfYear = newMonthOfYear
	return nil
}

func (monthOfYear MonthOfYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(monthOfYear.String())
}
