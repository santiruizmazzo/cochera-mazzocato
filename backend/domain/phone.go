package domain

import (
	"errors"
	"strconv"
	"strings"
)

type Phone struct {
	CountryCode string
	LineNumber  string
}

var (
	ErrPhoneMustStartWithPlusSign  = errors.New("el teléfono debe comenzar con un símbolo de +")
	ErrPhoneTooLong                = errors.New("el teléfono debe tener 15 dígitos como máximo")
	ErrPhoneFullOfZeroes           = errors.New("el teléfono no puede estar lleno de ceros")
	ErrPhoneMustContainNumbersOnly = errors.New("el teléfono solo puede tener números")
	ErrPhoneMustBeString           = errors.New("el teléfono debe ser un string")
)

func NewPhone(rawValue any) (Phone, error) {
	var phone Phone
	var runes []rune

	switch value := rawValue.(type) {
	case string:
		runes = []rune(value)
	case nil:
		return Phone{CountryCode: "", LineNumber: ""}, nil
	default:
		return Phone{}, ErrPhoneMustBeString
	}

	if string(runes[0:1]) != "+" {
		return Phone{}, ErrPhoneMustStartWithPlusSign
	}

	if len(runes) > 15 {
		return Phone{}, ErrPhoneTooLong
	}

	substring := strings.TrimPrefix(string(runes), "+")
	integerPhone, err := strconv.Atoi(substring)

	if err != nil {
		return Phone{}, ErrPhoneMustContainNumbersOnly
	}

	if integerPhone == 0 {
		return Phone{}, ErrPhoneFullOfZeroes
	}

	if len(runes[1:]) == 12 && string(runes[1:3]) == "54" {
		phone = Phone{
			CountryCode: "54",
			LineNumber:  string(runes[3:]),
		}
	}

	if len(runes[1:]) == 11 && string(runes[1:4]) == "598" {
		phone = Phone{
			CountryCode: "598",
			LineNumber:  string(runes[4:]),
		}
	}

	return phone, nil
}
