package value

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewEmail_ReturnsValidEmail(t *testing.T) {
	emails := []string{
		strings.Repeat("a",248)+"@b.com",
		"email@example.com",
		"firstname.lastname@example.com",
		"email@subdomain.example.com",
		"firstname+lastname@example.com",
		"email@123.123.123.123",
		"1234567890@example.com",
		"email@example-one.com",
		"_______@example.com",
		"email@example.name",
		"email@example.museum",
		"email@example.co.jp",
		"firstname-lastname@example.com",
	}

	for _, s := range emails {
		t.Run(s, func(t *testing.T) {
			email, err := NewEmail(s)

			assert.Nil(t, err)
			assert.Equal(t, s, email.Value())
		})
	}
}

func TestNewEmail_ReturnsErrorIfInvalidEmail(t *testing.T) {
	emails := []string{
		"a@a",
		strings.Repeat("a",249)+"@b.com",
		"email@[123.123.123.123]",
		"\"email\"@example.com",
		"plainaddress",
		"#@%^%#$@#$@#.com",
		"@example.com",
		"Joe Smith <email@example.com>",
		"email.example.com",
		"email@example@example.com",
		"あいうえお@example.com",
		"email@example.com (Joe Smith)",
		"email@example",
		"email@-example.com",
		"email@example..com",
	}

	for _, s := range emails {

		t.Run(s, func(t *testing.T) {
			email, err := NewEmail(s)

			assert.NotNil(t, err)
			assert.Empty(t, email)
		})
	}
}

func TestMustNewEmail_ReturnsValidEmail(t *testing.T) {
	s := "email@example.com"

	email := MustNewEmail(s)

	assert.Equal(t, s, email.Value())
}

func TestMustNewEmail_PanicsIfInvalidEmail(t *testing.T) {
	s := "@example.com"

	assert.Panics(t, func() {
		MustNewEmail(s)
	})
}
