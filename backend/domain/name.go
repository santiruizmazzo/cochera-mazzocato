package domain

import "errors"

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
