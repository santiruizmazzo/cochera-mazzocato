package domain

import (
	"encoding/json"
	"errors"
)

type Address string

const MAX_LENGTH int = 100

var (
	ErrAddressMustBeString = errors.New("el domicilio debe ser un string")
	ErrAddressTooLong      = errors.New("el domicilio debe tener 100 caracteres como máximo")
)

func NewAddress(rawValue any) (Address, error) {
	var stringValue string

	switch value := rawValue.(type) {
	case string:
		stringValue = value
	case *string:
		if value == nil {
			return Address(""), nil
		}
		stringValue = *value
	case nil:
		return Address(""), nil
	default:
		return Address(""), ErrAddressMustBeString
	}

	if len(stringValue) > MAX_LENGTH {
		return Address(""), ErrAddressTooLong
	}
	return Address(stringValue), nil
}

func (address Address) IsEmpty() bool {
	return address == ""
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
	if address.IsEmpty() {
		return json.Marshal(nil)
	}
	return json.Marshal(string(address))
}
