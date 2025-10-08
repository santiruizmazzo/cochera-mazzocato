package domain

import "errors"

var (
	ErrDuplicateDNI           = errors.New("el DNI ya existe")
	ErrDuplicateEmail         = errors.New("el email ya está en uso")
	ErrTenantNotFound         = errors.New("inquilino no encontrado")
	ErrNoMatchingTenantsFound = errors.New("no se encontraron inquilinos que coincidan")
)
