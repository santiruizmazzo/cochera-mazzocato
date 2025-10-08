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
	ErrDNINotInValidRange = errors.New("el DNI debe ser un entero positivo")
	ErrDNIMustBeANumber   = errors.New("el DNI debe ser un número")
)

func NewDNI(rawValue any) (DNI, error) {
	var integerValue int

	switch value := rawValue.(type) {
	case int:
		integerValue = value
	case float64:
		integerValue = int(value)
	default:
		return DNI(0), ErrDNIMustBeANumber
	}

	if integerValue < DNI_MIN_VALUE || integerValue > DNI_MAX_VALUE {
		return DNI(0), ErrDNINotInValidRange
	}

	return DNI(integerValue), nil
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
