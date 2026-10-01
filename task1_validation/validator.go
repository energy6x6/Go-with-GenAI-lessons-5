// Package validation перевіряє форму реєстрації та повертає всі невалідні поля.
package validation

import (
	"strings"
	"unicode/utf8"
)

type RegistrationForm struct {
	Email    string
	Password string
	Age      int
}

// ValidationError містить назви всіх невалідних полів.
type ValidationError struct {
	Fields []string
}

func (e *ValidationError) Error() string {
	return "registration invalid: fields " + strings.Join(e.Fields, ", ")
}

// ValidateRegistration перевіряє непорожній email, пароль від 8 рун та вік [0, 150].
func ValidateRegistration(f RegistrationForm) error {
	var fields []string
	if f.Email == "" {
		fields = append(fields, "email")
	}
	if utf8.RuneCountInString(f.Password) < 8 {
		fields = append(fields, "password")
	}
	if f.Age < 0 || f.Age > 150 {
		fields = append(fields, "age")
	}
	if len(fields) > 0 {
		return &ValidationError{Fields: fields}
	}
	return nil
}
