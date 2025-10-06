package domain

type DNI struct {
	value uint32
}

func NewDNI(number uint32) (DNI, error) {
	return DNI{value: number}, nil
}
