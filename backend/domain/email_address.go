package domain

import (
	"encoding/json"
	"errors"
	"net/mail"
)

type EmailAddress string

var (
	ErrInvalidEmailFormat = errors.New("el email debe seguir el formato estándar")
	ErrEmailTooLong       = errors.New("el email debe tener 100 caracteres como máximo")
	ErrEmailMustBeString  = errors.New("el email debe ser un string")
)

func NewEmailAddress(rawValue any) (EmailAddress, error) {
	var validEmailAddress *mail.Address
	var err error

	switch value := rawValue.(type) {
	case string:
		validEmailAddress, err = mail.ParseAddress(value)
	case nil:
		return EmailAddress(""), nil
	case *string:
		if value == nil {
			return EmailAddress(""), nil
		}
		validEmailAddress, err = mail.ParseAddress(*value)
	default:
		return EmailAddress(""), ErrEmailMustBeString
	}

	if err != nil {
		return EmailAddress(""), ErrInvalidEmailFormat
	}

	if len(validEmailAddress.Address) > 100 {
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
