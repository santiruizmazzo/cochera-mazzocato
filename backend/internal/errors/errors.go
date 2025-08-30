package myerrors

import "errors"

var (
	ErrDuplicateDNI   = errors.New("dni already exists")
	ErrDuplicateEmail = errors.New("email already exists")
)
