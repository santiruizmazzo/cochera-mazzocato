package domain

import (
	"encoding/json"
	"errors"
)

type Address string

var (
	ErrAddressMustBeString = errors.New("el domicilio debe ser un string")
	ErrAddressTooLong      = errors.New("el domicilio debe tener 100 caracteres como máximo")
)

func NewAddress(rawValue any) (Address, error) {
	switch value := rawValue.(type) {
	case string:
		if len(value) > 100 {
			return Address(""), ErrAddressTooLong
		}
		return Address(value), nil
	case nil:
		return Address(""), nil
	}
	return Address(""), ErrAddressMustBeString
}

func (address *Address) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validAddress, err := NewAddress(rawValue)
	if err != nil {
		return err
	}

	*address = validAddress
	return nil
}

func (address Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(address))
}
