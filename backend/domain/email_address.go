package domain

import (
	"errors"
	"net/mail"
)

type EmailAddress struct {
	Value string
}

var (
	ErrInvalidEmailFormat = errors.New("el email debe seguir el formato estándar")
	ErrEmailTooLong       = errors.New("el email debe tener 100 caracteres como máximo")
)

func NewEmailAddress(rawValue any) (EmailAddress, error) {
	var validEmailAddress *mail.Address
	var err error

	switch value := rawValue.(type) {
	case string:
		validEmailAddress, err = mail.ParseAddress(value)
	case nil:
		return EmailAddress{Value: ""}, nil
	}

	if err != nil {
		return EmailAddress{}, ErrInvalidEmailFormat
	}

	if len(validEmailAddress.String()) > 100 {
		return EmailAddress{}, ErrEmailTooLong
	}

	return EmailAddress{Value: validEmailAddress.String()}, nil
}
