package time

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MonthOfYear struct {
	month uint8
	year  uint16
}

func NewMonthOfYear(month uint8, year uint16) MonthOfYear {
	return MonthOfYear{month: month, year: year}
}

func (monthOfYear MonthOfYear) String() string {
	return fmt.Sprintf("%02d-%04d", monthOfYear.month, monthOfYear.year)
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

var ErrMonthNotParseable = errors.New("imposible to parse this month")
var ErrYearNotParseable = errors.New("imposible to parse this year")

func NewMonthOfYearFromString(monthOfYearString string) (MonthOfYear, error) {
	splitString := strings.Split(monthOfYearString, "-")
	month, err := StringToUint8(splitString[0])
	if err != nil {
		return MonthOfYear{}, ErrMonthNotParseable
	}

	year, err := StringToUint16(splitString[1])
	if err != nil {
		return MonthOfYear{}, ErrYearNotParseable
	}

	return MonthOfYear{month: month, year: year}, nil
}

func (monthOfYear *MonthOfYear) MarshalJSON() ([]byte, error) {
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
