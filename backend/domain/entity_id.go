package domain

import "errors"

type EntityID uint32

var ErrIDMustBeANumber = errors.New("el ID debe ser un número")

func NewEntityID(rawValue any) (EntityID, error) {
	switch value := rawValue.(type) {
	case int:
		return EntityID(value), nil
	default:
		return EntityID(0), ErrIDMustBeANumber
	}
}
