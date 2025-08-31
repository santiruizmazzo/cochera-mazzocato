package myerrors

import "errors"

var (
	ErrRequiredDNI                 = errors.New("dni is required")
	ErrDNIMustBeNumber             = errors.New("dni must be a positive integer")
	ErrDuplicateDNI                = errors.New("dni already exists")
	ErrRequiredName                = errors.New("name is required")
	ErrNameMustBeString            = errors.New("name must be a string")
	ErrInvalidName                 = errors.New("name must be a string")
	ErrDuplicateEmail              = errors.New("email already exists")
	ErrInvalidLastName             = errors.New("last name must be a string")
	ErrInvalidAddress              = errors.New("address must be a string")
	ErrInvalidPhone                = errors.New("phone must be a string")
	ErrInvalidEmail                = errors.New("email must be a string")
	ErrRequiredEntryMonth          = errors.New("entry month is required")
	ErrEntryMonthMustBeString      = errors.New("entry month must be string")
	ErrRequiredLastName            = errors.New("last name is required")
	ErrLastNameMustBeString        = errors.New("last name must be a string")
	ErrRequiredAddress             = errors.New("address is required")
	ErrAddressMustBeString         = errors.New("address must be a string")
	ErrRequiredPhone               = errors.New("phone is required")
	ErrPhoneMustBeString           = errors.New("phone must be a string")
	ErrRequiredEmail               = errors.New("email is required")
	ErrEmailMustBeString           = errors.New("email must be a string")
	ErrPhoneMustStartWithPlusSign  = errors.New("phone must start with + sign")
	ErrPhoneMustContainNumbersOnly = errors.New("phone must contain numbers only")
)
