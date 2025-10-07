package domain

import (
	"errors"
	"math"
)

type EntityID uint32

var (
	ErrIDMustBeAnInteger = errors.New("el ID debe ser un número entero")
	ErrIDTooSmall        = errors.New("el ID debe ser un entero mayor o igual a 1")
	ErrIDTooBig          = errors.New("el ID debe ser un entero menor o igual a 4294967295")
)

func NewEntityID(rawValue any) (EntityID, error) {
	switch value := rawValue.(type) {
	case int:
		if value < 1 {
			return EntityID(0), ErrIDTooSmall
		}
		if value > math.MaxUint32 {
			return EntityID(0), ErrIDTooBig
		}
		return EntityID(value), nil
	default:
		return EntityID(0), ErrIDMustBeAnInteger
	}
}
