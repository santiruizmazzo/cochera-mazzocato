package domain

import "errors"

var (
	ErrDuplicateDNI   = errors.New("el DNI ya existe")
	ErrDuplicateEmail = errors.New("el email ya está en uso")
)
