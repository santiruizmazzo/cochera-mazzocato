package domain

type Address struct {
	Value string
}

func NewAddress(rawValue string) (Address, error) {
	return Address{Value: rawValue}, nil
}
