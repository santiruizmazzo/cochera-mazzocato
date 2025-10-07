package domain

type Address struct {
	Value string
}

func NewAddress(rawValue any) (Address, error) {
	switch value := rawValue.(type) {
	case string:
		return Address{Value: value}, nil
	case nil:
		return Address{}, nil
	}
	return Address{}, nil
}
