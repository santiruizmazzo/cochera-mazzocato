package domain

import (
	"errors"
	"math"
)

type DNI struct {
	Value uint32
}

var ErrDNINotInValidRange = errors.New("el DNI debe ser un entero positivo")

func NewDNI(rawNumber any) (DNI, error) {
	var validValue uint32

	switch value := rawNumber.(type) {
	case float64:
		if value < 1 || value > math.MaxUint32 {
			return DNI{}, ErrDNINotInValidRange
		}
		validValue = uint32(value)
	case int:
		if value < 1 || value > math.MaxUint32 {
			return DNI{}, ErrDNINotInValidRange
		}
		validValue = uint32(value)
	default:
		return DNI{}, nil
	}

	return DNI{Value: validValue}, nil
}
