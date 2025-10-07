package domain

import (
	"encoding/json"
	"errors"
)

type Name struct {
	Value string
}

var (
	ErrNameMustBeAString = errors.New("el nombre/apellido debe ser un string")
	ErrNameTooLong       = errors.New("el nombre debe tener 50 caracteres como máximo")
)

func NewName(rawValue any) (Name, error) {
	switch value := rawValue.(type) {
	case string:
		if len(value) > 50 {
			return Name{}, ErrNameTooLong
		}
		return Name{Value: value}, nil
	}
	return Name{}, ErrNameMustBeAString
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
