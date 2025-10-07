package domain

import "errors"

type Address struct {
	Value string
}

var (
	ErrAddressMustBeString = errors.New("el domicilio debe ser un string")
	ErrAddressTooLong      = errors.New("el domicilio debe tener 100 caracteres como máximo")
)

func NewAddress(rawValue any) (Address, error) {
	switch value := rawValue.(type) {
	case string:
		if len(value) > 100 {
			return Address{}, ErrAddressTooLong
		}
		return Address{Value: value}, nil
	case nil:
		return Address{}, nil
	}
	return Address{}, ErrAddressMustBeString
}
