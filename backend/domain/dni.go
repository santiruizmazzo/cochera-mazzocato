package domain

import (
	"encoding/json"
	"errors"
	"math"
)

type DNI uint32

const DNI_MIN_VALUE int = 1
const DNI_MAX_VALUE int = math.MaxUint32

var (
	ErrDNINotInValidRange = errors.New("el DNI debe estar entre 1 y 4294967295")
	ErrDNIMustBeAnInteger = errors.New("el DNI debe ser un número entero")
)

func NewDNI(rawValue any) (DNI, error) {
	integerDni, err := extractIntegerDNI(rawValue)
	if err != nil {
		return DNI(0), err
	}

	if !isInValidRange(integerDni) {
		return DNI(0), ErrDNINotInValidRange
	}

	return DNI(integerDni), nil
}

func extractIntegerDNI(rawValue any) (int, error) {
	switch value := rawValue.(type) {
	case int:
		return value, nil
	case float64:
		if value != math.Floor(value) {
			return 0, ErrDNIMustBeAnInteger
		}
		return int(value), nil
	default:
		return 0, ErrDNIMustBeAnInteger
	}
}

func isInValidRange(integerDni int) bool {
	return DNI_MIN_VALUE <= integerDni && integerDni <= DNI_MAX_VALUE
}

func (dni *DNI) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validDni, err := NewDNI(rawValue)
	if err != nil {
		return err
	}

	*dni = validDni
	return nil
}

func (dni DNI) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint32(dni))
}
