package validation

import (
	"net"
	"regexp"
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

    minClientIDLength = 1
    maxClientIDLength = 23

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

func ValidateClientID(clientID string) *errors.ErrorWrapper {
	if len(clientID) < minClientIDLength || len(clientID) > maxClientIDLength {
		return errors.ErrLengthNotInRange.Format(minClientIDLength, maxClientIDLength)
	}

    validID := regexp.MustCompile(`^[A-Za-z0-9\-_\.]+$`).MatchString(clientID)
    if !validID {
        return errors.ErrForbiddenCharacter.Format()
    }

    return nil
}

func ValidateLogin(login string) *errors.ErrorWrapper {
    if !lengthInRange(login, minLoginLength, maxLoginLength) {
        return errors.ErrLengthNotInRange.Format(minLoginLength, maxLoginLength)
    }

	if strings.ContainsRune(login, ' ') {
        return errors.ErrForbiddenCharacter
    }

    return nil
}

func ValidatePassword(password string) *errors.ErrorWrapper {
    if !lengthInRange(password, minPasswordLength, maxPasswordLength) {
        return errors.ErrLengthNotInRange.Format(minPasswordLength, maxPasswordLength)
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
        return errors.ErrLengthNotInRange.Format(minUsernameLength, maxUsernameLength)
    }

    return nil
}

func ParseIPAddress(ip string) (net.IP, *errors.ErrorWrapper) {
    if parsedIP := net.ParseIP(ip); parsedIP == nil {
        return parsedIP, nil
    }

    return nil, errors.ErrInvalidIPAddress
}