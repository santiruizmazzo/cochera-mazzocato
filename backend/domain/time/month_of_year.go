package time

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MonthOfYear struct {
	month uint8
	year  uint16
}

func (monthOfYear MonthOfYear) String() string {
	return fmt.Sprintf("%d-%d", monthOfYear.month, monthOfYear.year)
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

var ErrMonthNotParseable = errors.New("mes introducido imposible de parsear")
var ErrYearNotParseable = errors.New("año introducido imposible de parsear")

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

func (monthOfYear *MonthOfYear) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	log.Println("HASTA ACA LLEGO!!!")

	newMonthOfYear, err := NewMonthOfYearFromString(s)
	if err != nil {
		return err
	}

	*monthOfYear = newMonthOfYear
	return nil
}
