package domain

import (
	"errors"
	"net/mail"
)

type EmailAddress struct {
	Value string
}

var ErrInvalidEmailFormat = errors.New("el email debe seguir el formato estándar")

func NewEmailAddress(rawValue string) (EmailAddress, error) {
	value, err := mail.ParseAddress(rawValue)
	if err != nil {
		return EmailAddress{}, ErrInvalidEmailFormat
	}

	return EmailAddress{Value: value.String()}, nil
}
