package domain

import "errors"

type Phone struct {
	CountryCode string
	LineNumber  string
}

var ErrPhoneMustStartWithPlusSign = errors.New("el teléfono debe comenzar con un símbolo de +")

func NewPhone(rawValue string) (Phone, error) {
	runes := []rune(rawValue)
	if string(runes[0:1]) != "+" {
		return Phone{}, ErrPhoneMustStartWithPlusSign
	}

	return Phone{
		CountryCode: string(runes[1:3]),
		LineNumber:  string(runes[3:]),
	}, nil
}
