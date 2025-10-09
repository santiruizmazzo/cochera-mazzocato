package vo

import (
	"encoding/json"
	"errors"
)

type Address string

const ADDRESS_MAX_LENGTH int = 100

var (
	ErrAddressMustBeString = errors.New("el domicilio debe ser un string")
	ErrAddressTooLong      = errors.New("el domicilio debe tener 100 caracteres como máximo")
)

func NewAddress(rawValue any) (Address, error) {
	stringAddress, err := extractStringAddress(rawValue)
	if stringAddress == "" || err != nil {
		return Address(""), err
	}

	if len(stringAddress) > ADDRESS_MAX_LENGTH {
		return Address(""), ErrAddressTooLong
	}

	return Address(stringAddress), nil
}

func extractStringAddress(rawValue any) (string, error) {
	switch value := rawValue.(type) {
	case string:
		return value, nil
	case *string:
		if value == nil {
			return "", nil
		}
		return *value, nil
	case nil:
		return "", nil
	default:
		return "", ErrAddressMustBeString
	}
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
