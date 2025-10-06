package domain

type DNI struct {
	Value uint32
}

func NewDNI(rawNumber any) (DNI, error) {
	var validValue uint32

	switch value := rawNumber.(type) {
	case float64:
		validValue = uint32(value)
	case int:
		validValue = uint32(value)
	default:
		return DNI{}, nil
	}

	if validValue < 1 {
		return DNI{}, ErrDNIMustBeNumber
	}

	return DNI{Value: validValue}, nil
}
