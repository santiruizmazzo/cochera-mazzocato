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
	integerID, err := extractIntegerID(rawValue)
	if err != nil {
		return EntityID(0), err
	}

	return createValidEntityID(integerID)
}

func extractIntegerID(rawValue any) (int, error) {
	switch value := rawValue.(type) {
	case int:
		return value, nil
	case float64:
		return int(value), nil
	default:
		return 0, ErrIDMustBeAnInteger
	}
}

func createValidEntityID(integerID int) (EntityID, error) {
	if integerID < ID_MIN_VALUE {
		return EntityID(0), ErrIDTooSmall
	}
	if integerID > ID_MAX_VALUE {
		return EntityID(0), ErrIDTooBig
	}
	return EntityID(integerID), nil
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
