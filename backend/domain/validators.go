package domain

import (
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

func extractAndValidateDNI(tenantMap map[string]any) (uint32, error) {
	raw, exists := tenantMap["dni"]
	if !exists {
		return 0, ErrRequiredDNI
	}

	return validateDNI(raw)
}

func validateDNI(rawDNI any) (uint32, error) {
	switch value := rawDNI.(type) {
	case float64: // float64 is how json.Unmarshal decodes every number by default
		if value < 1 || value > 4294967295 {
			return 0, ErrDNIMustBeNumber
		}
		return uint32(value), nil
	case int:
		return uint32(value), nil
	case uint32:
		return value, nil
	default:
		return 0, ErrDNIMustBeNumber
	}
}

func extractAndValidateName(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["name"]
	if !exists {
		return "", ErrRequiredName
	}

	return validateName(raw)
}

func validateName(rawName any) (string, error) {
	switch value := rawName.(type) {
	case string:
		if len(value) > 50 {
			return "", ErrNameTooLong
		}

		return value, nil
	default:
		return "", ErrNameMustBeString
	}
}

func extractAndValidateLastName(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["last_name"]
	if !exists {
		return "", ErrRequiredLastName
	}

	return validateLastName(raw)
}

func validateLastName(rawLastName any) (string, error) {
	switch value := rawLastName.(type) {
	case string:
		if len(value) > 50 {
			return "", ErrLastNameTooLong
		}

		return value, nil
	default:
		return "", ErrLastNameMustBeString
	}
}

func extractAndValidateAddress(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["address"]
	if !exists {
		return "", nil
	}

	return validateAddress(raw)
}

func validateAddress(rawAddress any) (string, error) {
	switch value := rawAddress.(type) {
	case string:
		if len(value) > 100 {
			return "", ErrAddressTooLong
		}

		return value, nil
	case nil:
		return "", nil
	default:
		return "", ErrAddressMustBeString
	}
}

func extractAndValidatePhone(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["phone"]
	if !exists {
		return "", nil
	}

	return validatePhone(raw)
}

func validatePhone(rawPhone any) (string, error) {
	switch value := rawPhone.(type) {
	case string:
		if value == "" {
			return value, nil
		}
		if !strings.HasPrefix(value, "+") {
			return "", ErrPhoneMustStartWithPlusSign
		}
		substring := strings.TrimPrefix(value, "+")
		if len(substring) > 15 {
			return "", ErrPhoneTooLong
		}

		integerPhone, err := strconv.Atoi(substring)
		if err != nil {
			return "", ErrPhoneMustContainNumbersOnly
		}
		if integerPhone == 0 {
			return "", ErrPhoneFullOfZeroes
		}
		return value, nil
	case nil:
		return "", nil
	default:
		return "", ErrPhoneMustBeString
	}
}

func extractAndValidateEmail(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["email"]
	if !exists {
		return "", nil
	}

	return validateEmail(raw)
}

func validateEmail(rawEmail any) (string, error) {
	switch value := rawEmail.(type) {
	case string:
		if value == "" {
			return value, nil
		}
		if _, err := mail.ParseAddress(value); err != nil {
			return "", ErrInvalidEmailFormat
		}
		if len(value) > 100 {
			return "", ErrEmailTooLong
		}
		return value, nil
	case nil:
		return "", nil
	default:
		return "", ErrEmailMustBeString
	}
}

func extractAndValidateEntryMonth(tenantMap map[string]any) (MonthOfYear, error) {
	raw, exists := tenantMap["entry_month"]
	if !exists {
		return MonthOfYear{}, ErrRequiredEntryMonth
	}

	return validateEntryMonth(raw)
}

func validateEntryMonth(rawEntryMonth any) (MonthOfYear, error) {
	switch value := rawEntryMonth.(type) {
	case MonthOfYear:
		return value, nil
	case string:
		matched, _ := regexp.MatchString(`^\d{2}-\d{4}$`, value)
		if !matched {
			return MonthOfYear{}, ErrEntryMonthInvalidFormat
		}

		return NewMonthOfYearFromString(value)
	default:
		return MonthOfYear{}, ErrEntryMonthMustBeString
	}
}
