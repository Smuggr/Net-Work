package errors

import (
	"errors"
	"fmt"
	"regexp"
)

type ErrorWrapper struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func NewErrorWrapper(key string, err error) *ErrorWrapper {
	return &ErrorWrapper{
		Key:     key,
		Message: err.Error(),
		Err:     err,
	}
}

func (e *ErrorWrapper) RemoveFormattingCharacters() {
	re := regexp.MustCompile(`%[a-zA-Z]`)
	e.Message = re.ReplaceAllString(e.Message, "")
}

func (e *ErrorWrapper) Error() string {
	e.RemoveFormattingCharacters()
	return e.Message
}

func (e *ErrorWrapper) FormatError(vars ...interface{}) string {
	if vars == nil {
		return e.Error()
    }

	return fmt.Sprintf(e.Message, vars...)
}

func (e *ErrorWrapper) Format(vars ...interface{}) *ErrorWrapper {
	e.Message = fmt.Sprintf(e.Message, vars...)
	return e
}

var (
	// Authentication errors
	ErrInvalidCredentials     = NewErrorWrapper("ErrInvalidCredentials", errors.New("invalid credentials"))
	ErrUnauthorized           = NewErrorWrapper("ErrUnauthorized", errors.New("unauthorized"))

	// Token errors
	ErrCreatingToken          = NewErrorWrapper("ErrCreatingToken", errors.New("error creating token"))
	ErrSigningToken           = NewErrorWrapper("ErrSigningToken", errors.New("error signing token"))
	ErrInvalidToken           = NewErrorWrapper("ErrInvalidToken", errors.New("invalid token"))
	ErrInvalidTokenFormat     = NewErrorWrapper("ErrInvalidTokenFormat", errors.New("invalid token format"))

	// User errors
	ErrDefaultAdminNotFound   = NewErrorWrapper("ErrDefaultAdminNotFound", errors.New("default admin not found"))
	ErrUserAlreadyExists      = NewErrorWrapper("ErrUserAlreadyExists", errors.New("user already exists"))
	ErrUserNotFound           = NewErrorWrapper("ErrUserNotFound", errors.New("user not found"))
	ErrFetchingUsersFromDB    = NewErrorWrapper("ErrFetchingUsersFromDB", errors.New("error fetching users from database"))

	// Device errors
	ErrDeviceAlreadyExists    = NewErrorWrapper("ErrDeviceAlreadyExists", errors.New("device already exists"))
	ErrRegisteringDeviceInDB  = NewErrorWrapper("ErrRegisteringDeviceInDB", errors.New("error registering device in database"))
	ErrRemovingDeviceFromDB   = NewErrorWrapper("ErrRemovingDeviceFromDB", errors.New("error removing device from database"))
	ErrFetchingDevicesFromDB  = NewErrorWrapper("ErrFetchingDevicesFromDB", errors.New("error fetching devices from database"))
	ErrDeviceNotFound         = NewErrorWrapper("ErrDeviceNotFound", errors.New("device not found"))
	ErrInvalidIPAddress       = NewErrorWrapper("ErrInvalidIPAddress", errors.New("invalid ip address"))

	// Database errors
	ErrConnectingToDB         = NewErrorWrapper("ErrConnectingToDB", errors.New("error connecting to database"))
	ErrGettingDBConnection    = NewErrorWrapper("ErrGettingDBConnection", errors.New("error getting database connection"))
	ErrClosingDBConnection    = NewErrorWrapper("ErrClosingDBConnection", errors.New("error closing database connection"))
	ErrRegisteringUserInDB    = NewErrorWrapper("ErrRegisteringUserInDB", errors.New("error registering user in database"))
	ErrUpdatingUserInDB       = NewErrorWrapper("ErrUpdatingUserInDB", errors.New("error updating user in database"))
	ErrRemovingUserFromDB     = NewErrorWrapper("ErrRemovingUserFromDB", errors.New("error removing user from database"))

	// Validation errors
	ErrInsufficientCharacters = NewErrorWrapper("ErrInsufficientCharacters", errors.New("insufficient characters"))
	ErrLengthNotInRange       = NewErrorWrapper("ErrNotInRange", errors.New("length not in range"))
	ErrForbiddenCharacter     = NewErrorWrapper("ErrForbiddenCharacter", errors.New("forbidden character"))

	// File errors
	ErrReadingConfigFile      = NewErrorWrapper("ErrReadingConfigFile", errors.New("error reading config file"))
	ErrFormattingConfigFile   = NewErrorWrapper("ErrFormattingConfigFile", errors.New("error formatting config file"))
	ErrReadingEnvFile         = NewErrorWrapper("ErrReadingEnvFile", errors.New("error reading env file"))

	// General errors
	ErrHashingPassword        = NewErrorWrapper("ErrHashingPassword", errors.New("error hashing password"))
	ErrInvalidRequestPayload  = NewErrorWrapper("ErrInvalidRequestPayload", errors.New("invalid request payload"))

	// Uncategorized errors
	ErrTest                   = NewErrorWrapper("ErrTest", errors.New("error test %s, %d, %f, %t"))
)

