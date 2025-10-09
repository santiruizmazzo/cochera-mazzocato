package domain

import (
	"encoding/json"
	"errors"
	"net/mail"
)

type EmailAddress string

const EMAIL_MAX_LENGTH int = 100

var (
	ErrInvalidEmailFormat = errors.New("el email debe seguir el formato estándar")
	ErrEmailTooLong       = errors.New("el email debe tener 100 caracteres como máximo")
	ErrEmailMustBeString  = errors.New("el email debe ser un string")
)

func NewEmailAddress(rawValue any) (EmailAddress, error) {
	stringEmailAddress, err := extractStringEmailAddress(rawValue)
	if stringEmailAddress == "" || err != nil {
		return EmailAddress(""), err
	}

	return createValidEmailAddress(stringEmailAddress)
}

func extractStringEmailAddress(rawValue any) (string, error) {
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
		return "", ErrEmailMustBeString
	}
}

func createValidEmailAddress(stringEmailAddress string) (EmailAddress, error) {
	validEmailAddress, err := mail.ParseAddress(stringEmailAddress)
	if err != nil {
		return EmailAddress(""), ErrInvalidEmailFormat
	}

	if len(validEmailAddress.Address) > EMAIL_MAX_LENGTH {
		return EmailAddress(""), ErrEmailTooLong
	}

	return EmailAddress(validEmailAddress.Address), nil
}

func (emailAddress *EmailAddress) UnmarshalJSON(data []byte) error {
	var rawValue any
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	validEmailAddress, err := NewEmailAddress(rawValue)
	if err != nil {
		return err
	}

	*emailAddress = validEmailAddress
	return nil
}

func (emailAddress EmailAddress) MarshalJSON() ([]byte, error) {
	if emailAddress == "" {
		return json.Marshal(nil)
	}
	return json.Marshal(string(emailAddress))
}

func (emailAddress EmailAddress) IsEmpty() bool {
	return emailAddress == ""
}

func (emailAddress EmailAddress) String() string {
	return string(emailAddress)
}
