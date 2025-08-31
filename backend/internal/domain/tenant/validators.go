package tenant

import (
	"cochera/internal/domain/calendar"
	myerrors "cochera/internal/errors"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

func extractAndValidateDNI(tenantMap map[string]any) (uint32, error) {
	raw, exists := tenantMap["dni"]
	if !exists {
		return 0, myerrors.ErrRequiredDNI
	}

	return validateDNI(raw)
}

func validateDNI(rawDNI any) (uint32, error) {
	switch value := rawDNI.(type) {
	case float64: // float64 is how json.Unmarshal decodes every number by default
		if value < 1 || value > 4294967295 {
			return 0, myerrors.ErrDNIMustBeNumber
		}
		return uint32(value), nil
	case int:
		return uint32(value), nil
	case uint32:
		return value, nil
	default:
		return 0, myerrors.ErrDNIMustBeNumber
	}
}

func extractAndValidateName(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["name"]
	if !exists {
		return "", myerrors.ErrRequiredName
	}

	return validateName(raw)
}

func validateName(rawName any) (string, error) {
	switch value := rawName.(type) {
	case string:
		if len(value) > 50 {
			return "", myerrors.ErrNameTooLong
		}

		return value, nil
	default:
		return "", myerrors.ErrNameMustBeString
	}
}

func extractAndValidateLastName(tenantMap map[string]any) (string, error) {
	raw, exists := tenantMap["last_name"]
	if !exists {
		return "", myerrors.ErrRequiredLastName
	}

	return validateLastName(raw)
}

func validateLastName(rawLastName any) (string, error) {
	switch value := rawLastName.(type) {
	case string:
		if len(value) > 50 {
			return "", myerrors.ErrLastNameTooLong
		}

		return value, nil
	default:
		return "", myerrors.ErrLastNameMustBeString
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
			return "", myerrors.ErrAddressTooLong
		}

		return value, nil
	default:
		return "", myerrors.ErrAddressMustBeString
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
			return "", myerrors.ErrPhoneMustStartWithPlusSign
		}
		substring := strings.TrimPrefix(value, "+")
		if len(substring) > 15 {
			return "", myerrors.ErrPhoneTooLong
		}

		integerPhone, err := strconv.Atoi(substring)
		if err != nil {
			return "", myerrors.ErrPhoneMustContainNumbersOnly
		}
		if integerPhone == 0 {
			return "", myerrors.ErrPhoneFullOfZeroes
		}
		return value, nil
	default:
		return "", myerrors.ErrPhoneMustBeString
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
			return "", myerrors.ErrInvalidEmailFormat
		}
		if len(value) > 100 {
			return "", myerrors.ErrEmailTooLong
		}
		return value, nil
	default:
		return "", myerrors.ErrEmailMustBeString
	}
}

func extractAndValidateEntryMonth(tenantMap map[string]any) (calendar.MonthOfYear, error) {
	raw, exists := tenantMap["entry_month"]
	if !exists {
		return calendar.MonthOfYear{}, myerrors.ErrRequiredEntryMonth
	}

	return validateEntryMonth(raw)
}

func validateEntryMonth(rawEntryMonth any) (calendar.MonthOfYear, error) {
	switch value := rawEntryMonth.(type) {
	case calendar.MonthOfYear:
		return value, nil
	case string:
		matched, _ := regexp.MatchString(`^\d{2}-\d{4}$`, value)
		if !matched {
			return calendar.MonthOfYear{}, myerrors.ErrEntryMonthInvalidFormat
		}

		return calendar.NewMonthOfYearFromString(value)
	default:
		return calendar.MonthOfYear{}, myerrors.ErrEntryMonthMustBeString
	}
}
