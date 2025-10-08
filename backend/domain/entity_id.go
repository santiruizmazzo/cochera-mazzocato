package domain

import (
	"encoding/json"
	"errors"
	"math"
)

type EntityID uint32

const ID_MIN_VALUE int = 1
const ID_MAX_VALUE int = math.MaxUint32

var (
	ErrIDMustBeAnInteger = errors.New("el ID debe ser un número entero")
	ErrIDTooSmall        = errors.New("el ID debe ser un entero mayor o igual a 1")
	ErrIDTooBig          = errors.New("el ID debe ser un entero menor o igual a 4294967295")
)

func NewEntityID(rawValue any) (EntityID, error) {
	var integerValue int

	switch value := rawValue.(type) {
	case int:
		integerValue = value
	case float64:
		integerValue = int(value)
	default:
		return EntityID(0), ErrIDMustBeAnInteger
	}

	if integerValue < ID_MIN_VALUE {
		return EntityID(0), ErrIDTooSmall
	}
	if integerValue > ID_MAX_VALUE {
		return EntityID(0), ErrIDTooBig
	}
	return EntityID(integerValue), nil
}

func (id *EntityID) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validId, err := NewEntityID(rawValue)
	if err != nil {
		return err
	}

	*id = validId
	return nil
}

func (id EntityID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint32(id))
}
