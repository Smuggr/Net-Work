package errors

import (
	"errors"
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

func (e *ErrorWrapper) Error() string {
	return e.Message
}

var (
	ErrInsufficientLength    = NewErrorWrapper("ErrInsufficientLength", errors.New("insufficient uppercase letters"))
	ErrInsufficientUppercase = NewErrorWrapper("ErrInsufficientUppercase", errors.New("insufficient uppercase letters"))
	ErrInsufficientLowercase = NewErrorWrapper("ErrInsufficientLowercase", errors.New("insufficient lowercase letters"))
	ErrInsufficientDigits    = NewErrorWrapper("ErrInsufficientDigits", errors.New("insufficient digits"))
	ErrInsufficientSpecial   = NewErrorWrapper("ErrInsufficientSpecial", errors.New("insufficient special characters"))

	ErrHashingPassword       = NewErrorWrapper("ErrHashingPassword", errors.New("error hashing password"))
	ErrInvalidRequestPayload = NewErrorWrapper("ErrInvalidRequestPayload", errors.New("invalid request payload"))
	ErrInvalidCredentials    = NewErrorWrapper("ErrInvalidCredentials", errors.New("invalid credentials"))
	ErrCreatingToken         = NewErrorWrapper("ErrCreatingToken", errors.New("error creating token"))
	ErrSigningToken          = NewErrorWrapper("ErrSigningToken", errors.New("error signing token"))
	ErrInvalidToken          = NewErrorWrapper("ErrInvalidToken", errors.New("invalid token"))
	ErrInvalidTokenFormat    = NewErrorWrapper("ErrInvalidTokenFormat", errors.New("invalid token format"))
	ErrUnauthorized          = NewErrorWrapper("ErrUnauthorized", errors.New("unauthorized"))

	ErrConnectingToDB        = NewErrorWrapper("ErrConnectingToDB", errors.New("error connecting to database"))
	ErrGettingDBConnection   = NewErrorWrapper("ErrGettingDBConnection", errors.New("error getting database connection"))
	ErrClosingDBConnection   = NewErrorWrapper("ErrClosingDBConnection", errors.New("error closing database connection"))
	ErrRegisteringUserInDB   = NewErrorWrapper("ErrRegisteringUserInDB", errors.New("error registering user in database"))
	ErrUpdatingUserInDB      = NewErrorWrapper("ErrUpdatingUserInDB", errors.New("error updating user in database"))
	ErrDefaultAdminNotFound  = NewErrorWrapper("ErrDefaultAdminNotFound", errors.New("default admin not found"))
	ErrUserAlreadyExists     = NewErrorWrapper("ErrUserAlreadyExists", errors.New("user already exists"))
	ErrUserNotFound          = NewErrorWrapper("ErrUserNotFound", errors.New("user not found"))

	ErrReadingConfigFile     = NewErrorWrapper("ErrReadingConfigFile", errors.New("error reading config file"))
	ErrFormattingConfigFile  = NewErrorWrapper("ErrFormattingConfigFile", errors.New("error formatting config file"))
	ErrReadingEnvFile        = NewErrorWrapper("ErrReadingEnvFile", errors.New("error reading env file"))
)