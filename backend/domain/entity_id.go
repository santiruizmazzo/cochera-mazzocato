package domain

import "errors"

type EntityID uint32

var (
	ErrIDMustBeAnInteger = errors.New("el ID debe ser un número entero")
	ErrIDTooSmall        = errors.New("el ID debe ser un entero mayor o igual a 1")
)

func NewEntityID(rawValue any) (EntityID, error) {
	switch value := rawValue.(type) {
	case int:
		if value < 1 {
			return EntityID(0), ErrIDTooSmall
		}
		return EntityID(value), nil
	default:
		return EntityID(0), ErrIDMustBeAnInteger
	}
}
