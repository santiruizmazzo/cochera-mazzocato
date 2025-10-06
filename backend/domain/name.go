package domain

import "errors"

type Name struct {
	Value string
}

var ErrNameMustBeAString = errors.New("el nombre/apellido debe ser un string")

func NewName(rawValue any) (Name, error) {
	switch value := rawValue.(type) {
	case string:
		return Name{Value: value}, nil
	}
	return Name{}, ErrNameMustBeAString
}
