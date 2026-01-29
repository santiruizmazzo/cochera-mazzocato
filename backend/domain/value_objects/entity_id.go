package vo

import (
	"encoding/json"
	"errors"
	"math"
)

type EntityID int

const ID_MIN_VALUE int = 0
const ID_MAX_VALUE int = math.MaxUint32

var (
	ErrIDMustBeAnInteger = errors.New("el ID debe ser un número entero")
	ErrIDTooSmall        = errors.New("el ID debe ser un entero mayor o igual a 0")
	ErrIDTooBig          = errors.New("el ID debe ser un entero menor o igual a 4294967295")
)

func NewEntityID(rawValue any) (EntityID, error) {
	integerID, err := extractIntegerID(rawValue)
	if err != nil {
		return EntityID(0), err
	}

	entityID := EntityID(integerID)
	if err = entityID.Validate(); err != nil {
		return EntityID(0), err
	}

	return entityID, nil
}

func extractIntegerID(rawValue any) (int, error) {
	switch value := rawValue.(type) {
	case int:
		return value, nil
	case float64:
		return int(value), nil
	case nil:
		return 0, nil
	default:
		return 0, ErrIDMustBeAnInteger
	}
}

func (id EntityID) Validate() error {
	if int(id) < ID_MIN_VALUE {
		return ErrIDTooSmall
	}
	if int(id) > ID_MAX_VALUE {
		return ErrIDTooBig
	}
	return nil
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
	if int(id) == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(uint32(id))
}
