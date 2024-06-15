package validator

import (
	"net"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"smuggr.xyz/net-work/common/logger"
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

	minUppercase = 1
	minLowercase = 1
	minDigit     = 1
	minSpecial   = 1
)

var (
	reservedClientIDs = []string{"inline"}
	reservedLogins    = []string{"administrator"}
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

func ValidateClientID(clientID string) *logger.MessageWrapper {
	for _, reservedClientID := range reservedClientIDs {
		if clientID == reservedClientID {
			return logger.ErrResourceAlreadyExists.Format(clientID, logger.ResourceDevice)
		}
	}

	if len(clientID) < minClientIDLength || len(clientID) > maxClientIDLength {
		return logger.ErrLengthNotInRange.Format(minClientIDLength, maxClientIDLength)
	}

	validID := regexp.MustCompile(`^[A-Za-z0-9\-_\.]+$`).MatchString(clientID)
	if !validID {
		return logger.ErrForbiddenCharacter.Format()
	}

	return nil
}

func ValidateLogin(login string) *logger.MessageWrapper {
	for _, reservedLogin := range reservedLogins {
		if login == reservedLogin {
			return logger.ErrResourceAlreadyExists.Format(login, logger.ResourceUser)
		}
	}

	if !lengthInRange(login, minLoginLength, maxLoginLength) {
		return logger.ErrLengthNotInRange.Format(minLoginLength, maxLoginLength)
	}

	if strings.ContainsRune(login, ' ') {
		return logger.ErrForbiddenCharacter
	}

	return nil
}

func ValidatePassword(password string) *logger.MessageWrapper {
	if !lengthInRange(password, minPasswordLength, maxPasswordLength) {
		return logger.ErrLengthNotInRange.Format(minPasswordLength, maxPasswordLength)
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
			return logger.ErrInsufficientCharacters
		}
	}

	return nil
}

func ValidateUsername(username string) *logger.MessageWrapper {
	if !lengthInRange(username, minUsernameLength, maxUsernameLength) {
		return logger.ErrLengthNotInRange.Format(minUsernameLength, maxUsernameLength)
	}

	return nil
}

func ValidateDisplayName(username string) *logger.MessageWrapper {
	return ValidateUsername(username)
}

func ParseIPAddress(ip string) (net.IP, *logger.MessageWrapper) {
	if parsedIP := net.ParseIP(ip); parsedIP == nil {
		return parsedIP, nil
	}

	return nil, logger.ErrInvalidIPAddress
}
