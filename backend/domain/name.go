package domain

type Name struct {
	Value string
}

func NewName(value string) (Name, error) {
	return Name{Value: value}, nil
}
