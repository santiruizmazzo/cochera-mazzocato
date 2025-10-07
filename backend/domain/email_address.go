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

func NewEmailAddress(rawValue string) (EmailAddress, error) {
	value, err := mail.ParseAddress(rawValue)
	if err != nil {
		return EmailAddress{}, ErrInvalidEmailFormat
	}

	if len(value.String()) > 100 {
		return EmailAddress{}, ErrEmailTooLong
	}

	return EmailAddress{Value: value.String()}, nil
}
