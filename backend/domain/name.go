package domain

import (
	"encoding/json"
	"errors"
)

type Name string

const NAME_MAX_LENGTH int = 50

var (
	ErrNameMustBeString = errors.New("el nombre/apellido debe ser un string")
	ErrNameTooLong      = errors.New("el nombre/apellido debe tener 50 caracteres como máximo")
)

func NewName(rawValue any) (Name, error) {
	switch value := rawValue.(type) {
	case string:
		if len(value) > NAME_MAX_LENGTH {
			return Name(""), ErrNameTooLong
		}
		return Name(value), nil
	default:
		return Name(""), ErrNameMustBeString
	}
}

func (name *Name) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validName, err := NewName(rawValue)
	if err != nil {
		return err
	}

	*name = validName
	return nil
}

func (name Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(name))
}
