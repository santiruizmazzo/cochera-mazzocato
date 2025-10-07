package domain

import "errors"

type Address struct {
	Value string
}

var ErrAddressMustBeString = errors.New("el domicilio debe ser un string")

func NewAddress(rawValue any) (Address, error) {
	switch value := rawValue.(type) {
	case string:
		return Address{Value: value}, nil
	case nil:
		return Address{}, nil
	}
	return Address{}, ErrAddressMustBeString
}
