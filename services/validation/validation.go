package validation

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"network/data/errors"
)

const (
    maxUsernameLength = 32
    minUsernameLength = 8
	minLoginLength    = 8
	maxLoginLength    = 16
    maxPasswordLength = 32
    minPasswordLength = 8

	minUppercase      = 1
	minLowercase      = 1
	minDigit          = 1
	minSpecial        = 1
)

func lengthInRange(s string, minLength, maxLength int) bool {
    length := utf8.RuneCountInString(s)
    return length >= minLength && length <= maxLength
}

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

func ValidateLogin(login string) *errors.ErrorWrapper {
    if !lengthInRange(login, minLoginLength, maxLoginLength) {
        return errors.ErrLengthNotInRange
    }

	if strings.ContainsRune(login, ' ') {
        return errors.ErrForbiddenCharacter
    }

    return nil
}

func ValidatePassword(password string) *errors.ErrorWrapper {
    if !lengthInRange(password, minPasswordLength, maxPasswordLength) {
        return errors.ErrLengthNotInRange
    }

    minCategories := []struct {
        category func(rune) bool
        minCount int
    }{
        {unicode.IsUpper, minUppercase},
        {unicode.IsLower, minLowercase},
        {unicode.IsDigit, minDigit},
        {func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }, minSpecial},
    }

    for _, c := range minCategories {
        if !containsCategory(password, c.category, c.minCount) {
            return errors.ErrInsufficientCharacters
        }
    }

    return nil
}

func ValidateUsername(username string) *errors.ErrorWrapper {
    if !lengthInRange(username, minUsernameLength, maxUsernameLength) {
        return errors.ErrLengthNotInRange
    }

    return nil
}