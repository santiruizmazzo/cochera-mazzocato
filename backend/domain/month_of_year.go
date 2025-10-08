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

const SEPARATION_CHARACTER string = "-"
const MONTH_MIN_VALUE uint64 = 1
const MONTH_MAX_VALUE uint64 = 12
const YEAR_MIN_VALUE uint64 = 0
const YEAR_MAX_VALUE uint64 = 9999

var (
	ErrMonthOfYearInvalidFormat = errors.New(`un mes de año debe seguir el formato: "MM-YYYY"`)
	ErrMonthOfYearMustBeString  = errors.New(`un mes de año debe ser un string con formato: "MM-YYYY"`)
	ErrMonthNotParseable        = errors.New("imposible procesar este mes")
	ErrMonthOfYearInvalidMonth  = errors.New("el mes debe estar entre 1 y 12")
	ErrYearNotParseable         = errors.New("imposible procesar este año")
	ErrMonthOfYearInvalidYear   = errors.New("el año debe estar entre 0 y 9999")
)

func NewMonthOfYear(rawValue any) (MonthOfYear, error) {
	var stringValue string
	var ok bool

	if stringValue, ok = rawValue.(string); !ok {
		return MonthOfYear{}, ErrMonthOfYearMustBeString
	}

	parts := strings.Split(stringValue, SEPARATION_CHARACTER)
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

func ParseMonth(stringMonth string) (uint8, error) {
	month, err := strconv.ParseUint(stringMonth, 10, 8)
	if err != nil {
		err = ErrMonthNotParseable
	}

	if month < MONTH_MIN_VALUE || month > MONTH_MAX_VALUE {
		return 0, ErrMonthOfYearInvalidMonth
	}

	return uint8(month), err
}

func ParseYear(stringYear string) (uint16, error) {
	year, err := strconv.ParseUint(stringYear, 10, 16)
	if err != nil {
		err = ErrYearNotParseable
	}

	if year < YEAR_MIN_VALUE || year > YEAR_MAX_VALUE {
		return 0, ErrMonthOfYearInvalidYear
	}

	return uint16(year), err
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
