package value

import (
	"errors"
	"regexp"
)

const emailPattern = "^" +
	"[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+" +
	"@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" +
	"(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+" +
	"$"
const emailInvalid = ""

var emailRegex = regexp.MustCompile(emailPattern)

type Email string

func NewEmail(s string) (Email, error) {
	if !isValidEmail(s) {
		return emailInvalid, errors.New("email is invalid")
	}

	e := Email(s)
	return e, nil
}

func MustNewEmail(s string) Email {
	e, err := NewEmail(s)
	if err != nil {
		panic(err)
	}

	return e
}

func (e Email) Value() string  {
	return string(e)
}

func isValidEmail(s string) bool {
	if len(s) < 3 || len(s) > 254 {
		return false
	}

	return emailRegex.MatchString(s)
}
