package vo

import (
	"encoding/json"
	"errors"
	"unicode"
)

type Name string

const NAME_MAX_LENGTH int = 50

var (
	ErrNameMustBeString              = errors.New("el nombre/apellido debe ser un string")
	ErrNameTooLong                   = errors.New("el nombre/apellido debe tener 50 caracteres como máximo")
	ErrNameMustIncludeCharactersOnly = errors.New("el nombre/apellido no puede contener números")
)

func NewName(rawValue any) (Name, error) {
	stringValue, ok := rawValue.(string)
	if !ok {
		return Name(""), ErrNameMustBeString
	}

	name := Name(stringValue)
	if err := name.Validate(); err != nil {
		return Name(""), err
	}

	return name, nil
}

func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
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

func (name Name) Validate() error {
	if len(name) > NAME_MAX_LENGTH {
		return ErrNameTooLong
	}

	if containsNumber(string(name)) {
		return ErrNameMustIncludeCharactersOnly
	}

	return nil
}
