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

func NewMonthOfYear(month uint8, year uint16) MonthOfYear {
	return MonthOfYear{Month: month, Year: year}
}

func (monthOfYear MonthOfYear) String() string {
	return fmt.Sprintf("%02d-%04d", monthOfYear.Month, monthOfYear.Year)
}

func StringToUint8(s string) (uint8, error) {
	newInt, err := StringToUint(s, 8)
	return uint8(newInt), err
}

func StringToUint16(s string) (uint16, error) {
	newInt, err := StringToUint(s, 16)
	return uint16(newInt), err
}

func StringToUint32(s string) (uint32, error) {
	newInt, err := StringToUint(s, 32)
	return uint32(newInt), err
}

func StringToUint(s string, maxBits int) (uint64, error) {
	n, err := strconv.ParseUint(s, 10, maxBits)
	if err != nil {
		return 0, err
	}
	return n, nil
}

var (
	ErrInvalidMonthOfYearString = errors.New(`un mes de año debe ser un string con formato: "MM-YYYY"`)
	ErrMonthNotParseable        = errors.New("imposible procesar este mes")
	ErrYearNotParseable         = errors.New("imposible procesar este año")
)

func NewMonthOfYearFromString(rawMonthOfYear any) (MonthOfYear, error) {
	var monthOfYearString string
	var ok bool

	if monthOfYearString, ok = rawMonthOfYear.(string); !ok {
		return MonthOfYear{}, ErrInvalidMonthOfYearString
	}

	splitString := strings.Split(monthOfYearString, "-")
	month, err := StringToUint8(splitString[0])
	if err != nil {
		return MonthOfYear{}, ErrMonthNotParseable
	}

	year, err := StringToUint16(splitString[1])
	if err != nil {
		return MonthOfYear{}, ErrYearNotParseable
	}

	return MonthOfYear{Month: month, Year: year}, nil
}

func (monthOfYear MonthOfYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(monthOfYear.String())
}

func (monthOfYear *MonthOfYear) UnmarshalJSON(data []byte) error {
	var stringFormatMonth string
	if err := json.Unmarshal(data, &stringFormatMonth); err != nil {
		return err
	}

	newMonthOfYear, err := NewMonthOfYearFromString(stringFormatMonth)
	if err != nil {
		return err
	}

	*monthOfYear = newMonthOfYear
	return nil
}
