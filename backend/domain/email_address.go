package domain

type EmailAddress struct {
	Value string
}

func NewEmailAddress(rawValue string) (EmailAddress, error) {
	return EmailAddress{Value: rawValue}, nil
}
