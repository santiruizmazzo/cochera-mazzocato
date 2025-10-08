package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Phone struct {
	CountryCode string
	LineNumber  string
}

const PHONE_START_SIGN string = "+"
const PHONE_MAX_LENGTH int = 15

const ARGENTINA_CODE string = "54"
const URUGUAY_CODE string = "598"

var lenghtByCountryCode map[string]int = map[string]int{
	ARGENTINA_CODE: 12,
	URUGUAY_CODE:   11,
}

var (
	ErrPhoneMustStartWithPlusSign  = errors.New("el teléfono debe comenzar con un símbolo de +")
	ErrPhoneTooLong                = errors.New("el teléfono debe tener 15 dígitos como máximo")
	ErrPhoneFullOfZeroes           = errors.New("el teléfono no puede estar lleno de ceros")
	ErrPhoneMustContainNumbersOnly = errors.New("el teléfono solo puede tener números")
	ErrPhoneMustBeString           = errors.New("el teléfono debe ser un string")
	ErrPhoneUnsupportedCountryCode = errors.New("el teléfono debe tener un prefijo aceptado: Argentina, Uruguay")
)

func NewPhone(rawValue any) (Phone, error) {
	stringNumber, err := extractStringNumber(rawValue)
	if stringNumber == "" || err != nil {
		return Phone{}, err
	}

	phoneNumber, err := extractValidNumber(stringNumber)
	if err != nil {
		return Phone{}, err
	}

	return createValidPhone(phoneNumber)
}

func extractStringNumber(rawValue any) (string, error) {
	switch value := rawValue.(type) {
	case string:
		return value, nil
	case *string:
		if value == nil {
			return "", nil
		}
		return *value, nil
	case nil:
		return "", nil
	default:
		return "", ErrPhoneMustBeString
	}
}

func extractValidNumber(phoneNumber string) (string, error) {
	if !strings.HasPrefix(phoneNumber, PHONE_START_SIGN) {
		return "", ErrPhoneMustStartWithPlusSign
	}

	if len(phoneNumber) > PHONE_MAX_LENGTH {
		return "", ErrPhoneTooLong
	}

	trimmedNumber := strings.TrimPrefix(phoneNumber, PHONE_START_SIGN)

	integerNumber, err := strconv.Atoi(trimmedNumber)
	if err != nil {
		return "", ErrPhoneMustContainNumbersOnly
	} else if integerNumber == 0 {
		return "", ErrPhoneFullOfZeroes
	}

	return trimmedNumber, nil
}

func createValidPhone(phoneNumber string) (Phone, error) {
	var phone Phone
	var err error

	switch true {
	case numberMatches(phoneNumber, ARGENTINA_CODE):
		phone.CountryCode = ARGENTINA_CODE
		phone.LineNumber = extractLineNumber(phoneNumber, ARGENTINA_CODE)
	case numberMatches(phoneNumber, URUGUAY_CODE):
		phone.CountryCode = URUGUAY_CODE
		phone.LineNumber = extractLineNumber(phoneNumber, URUGUAY_CODE)
	default:
		err = ErrPhoneUnsupportedCountryCode
	}

	return phone, err
}

func numberMatches(phoneNumber string, countryCode string) bool {
	return strings.HasPrefix(phoneNumber, countryCode) && len(phoneNumber) == lenghtByCountryCode[countryCode]
}

func extractLineNumber(phoneNumber string, countryCode string) string {
	return strings.TrimPrefix(phoneNumber, countryCode)
}

func (phone *Phone) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validPhone, err := NewPhone(rawValue)
	if err != nil {
		return err
	}

	*phone = validPhone
	return nil
}

func (phone Phone) MarshalJSON() ([]byte, error) {
	if phone.IsEmpty() {
		return json.Marshal(nil)
	}
	stringFormat := fmt.Sprintf("+%s%s", phone.CountryCode, phone.LineNumber)
	return json.Marshal(stringFormat)
}

func (phone Phone) IsEmpty() bool {
	return phone.CountryCode == "" && phone.LineNumber == ""
}

func (phone Phone) String() string {
	if phone.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("+%s%s", phone.CountryCode, phone.LineNumber)
}
