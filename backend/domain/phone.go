package domain

type Phone struct {
	CountryCode string
	LineNumber  string
}

func NewPhone(rawValue string) (Phone, error) {
	runes := []rune(rawValue)
	return Phone{
		CountryCode: string(runes[1:3]),
		LineNumber:  string(runes[3:]),
	}, nil
}
