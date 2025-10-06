package domain

type DNI struct {
	Value uint32
}

func NewDNI(number any) (DNI, error) {
	switch value := number.(type) {
	case float64:
		return DNI{Value: uint32(value)}, nil
	case int:
		return DNI{Value: uint32(value)}, nil
	default:
		return DNI{}, nil
	}
}
