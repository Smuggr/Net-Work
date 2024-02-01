package validation

import (
	"unicode"
	"unicode/utf8"

	"overseer/services/errors"
)


func containsCategory(s string, category func(rune) bool, n int) bool {
	count := 0
	for _, char := range s {
		if category(char) {
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}

func containsSpecial(s string, n int) bool {
	count := 0
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}

func ValidateLogin(login string) error {
	const (
		maxLength  = 32
		minLength  = 8
	)

	if length := utf8.RuneCountInString(login); length < minLength || length > maxLength {
		return errors.ErrInsufficientLength
	}

	return nil
}

func ValidatePassword(password string) error {
	const (
		maxLength  = 32
		minLength  = 8
		minUpper   = 1
		minLower   = 1
		minDigit   = 1
		minSpecial = 1
	)

	if length := utf8.RuneCountInString(password); length < minLength || length > maxLength {
		return errors.ErrInsufficientDigits
	}

	if !containsCategory(password, unicode.IsUpper, minUpper) {
		return errors.ErrInsufficientUppercase
	}

	if !containsCategory(password, unicode.IsLower, minLower) {
		return errors.ErrInsufficientLowercase
	}

	if !containsCategory(password, unicode.IsDigit, minDigit) {
		return errors.ErrInsufficientDigits
	}

	if !containsSpecial(password, minSpecial) {
		return errors.ErrInsufficientSpecial
	}

	return nil
}

func ValidateUsername(username string) error {
	const (
		maxLength  = 32
		minLength  = 8
	)

	if length := utf8.RuneCountInString(username); length < minLength || length > maxLength {
		return errors.ErrInsufficientLength
	}

	return nil
}