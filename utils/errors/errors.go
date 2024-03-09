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

func removeFormattingCharacters(message string) string {
	re := regexp.MustCompile(`%[a-zA-Z]`)
	return re.ReplaceAllString(message, "")
}

func NewErrorWrapper(key string, err error) *ErrorWrapper {
	return &ErrorWrapper{
		Key:     key,
		Message: err.Error(),
		Err:     err,
	}
}

func (e *ErrorWrapper) Error() string {
	return removeFormattingCharacters(e.Message)
}

func (e *ErrorWrapper) FormatError(vars ...interface{}) string {
	if vars == nil {
		return e.Error()
	}

	return fmt.Sprintf(e.Message, vars...)
}

func (e *ErrorWrapper) Format(vars ...interface{}) *ErrorWrapper {
	copy := *e
	copy.Message = e.FormatError(vars...)
	return &copy
}

var (
	// Authentication errors
	ErrInvalidCredentials = NewErrorWrapper("ErrInvalidCredentials", errors.New("invalid credentials"))
	ErrUnauthorized       = NewErrorWrapper("ErrUnauthorized", errors.New("unauthorized"))

	// Token errors
	ErrCreatingToken      = NewErrorWrapper("ErrCreatingToken", errors.New("error creating token"))
	ErrSigningToken       = NewErrorWrapper("ErrSigningToken", errors.New("error signing token"))
	ErrInvalidToken       = NewErrorWrapper("ErrInvalidToken", errors.New("invalid token"))
	ErrInvalidTokenFormat = NewErrorWrapper("ErrInvalidTokenFormat", errors.New("invalid token format"))

	// User errors
	ErrDefaultAdminNotFound = NewErrorWrapper("ErrDefaultAdminNotFound", errors.New("default admin not found"))
	ErrUserAlreadyExists    = NewErrorWrapper("ErrUserAlreadyExists", errors.New("user %s already exists"))
	ErrUserNotFound         = NewErrorWrapper("ErrUserNotFound", errors.New("user %s not found"))
	ErrFetchingUsersFromDB  = NewErrorWrapper("ErrFetchingUsersFromDB", errors.New("error fetching users from database"))

	// Device errors
	ErrDeviceAlreadyExists   = NewErrorWrapper("ErrDeviceAlreadyExists", errors.New("device %s already exists"))
	ErrRegisteringDeviceInDB = NewErrorWrapper("ErrRegisteringDeviceInDB", errors.New("error registering device %s in database"))
	ErrUpdatingDeviceInDB    = NewErrorWrapper("ErrUpdatingDeviceInDB", errors.New("error updating device %s in database"))
	ErrRemovingDeviceFromDB  = NewErrorWrapper("ErrRemovingDeviceFromDB", errors.New("error removing device %s from database"))
	ErrFetchingDevicesFromDB = NewErrorWrapper("ErrFetchingDevicesFromDB", errors.New("error fetching devices from database"))
	ErrDeviceNotFound        = NewErrorWrapper("ErrDeviceNotFound", errors.New("device %s not found"))
	ErrInvalidIPAddress      = NewErrorWrapper("ErrInvalidIPAddress", errors.New("invalid ip address"))
	ErrGettingPluginProvider = NewErrorWrapper("ErrGettingPluginProvider", errors.New("error getting plugin provider for plugin %s"))
	ErrPluginAlreadyLoaded   = NewErrorWrapper("ErrPluginAlreadyLoaded", errors.New("error plugin %s already loaded"))
	ErrPluginConflict        = NewErrorWrapper("ErrPluginConflict", errors.New("plugin conflict, %s already exists"))
	ErrAPIVersionMismatch    = NewErrorWrapper("ErrAPIVersionMismatch", errors.New("api version mismatch, plugin api version %s, expected %s"))
	ErrRemovingDevicePlugin  = NewErrorWrapper("ErrRemovingDevicePlugin", errors.New("error removing device %s plugin %s"))
	ErrCreatingDevicePlugin  = NewErrorWrapper("ErrCreatingDevicePlugin", errors.New("error creating device %s plugin %s"))
	ErrClientNotFound        = NewErrorWrapper("ErrClientNotFound", errors.New("client %s not found"))

	// Database errors
	ErrConnectingToDB      = NewErrorWrapper("ErrConnectingToDB", errors.New("error connecting to database"))
	ErrGettingDBConnection = NewErrorWrapper("ErrGettingDBConnection", errors.New("error getting database connection"))
	ErrClosingDBConnection = NewErrorWrapper("ErrClosingDBConnection", errors.New("error closing database connection"))
	ErrRegisteringUserInDB = NewErrorWrapper("ErrRegisteringUserInDB", errors.New("error registering user %s in database"))
	ErrUpdatingUserInDB    = NewErrorWrapper("ErrUpdatingUserInDB", errors.New("error updating user %s in database"))
	ErrRemovingUserFromDB  = NewErrorWrapper("ErrRemovingUserFromDB", errors.New("error removing user %s from database"))

	// Validation errors
	ErrInsufficientCharacters = NewErrorWrapper("ErrInsufficientCharacters", errors.New("insufficient characters"))
	ErrLengthNotInRange       = NewErrorWrapper("ErrLengthNotInRange", errors.New("length not in range <%d, %d>"))
	ErrForbiddenCharacter     = NewErrorWrapper("ErrForbiddenCharacter", errors.New("forbidden character(s)"))

	// File errors
	ErrReadingConfigFile     = NewErrorWrapper("ErrReadingConfigFile", errors.New("error reading config file"))
	ErrFormattingConfigFile  = NewErrorWrapper("ErrFormattingConfigFile", errors.New("error formatting config file"))
	ErrReadingEnvFile        = NewErrorWrapper("ErrReadingEnvFile", errors.New("error reading env file"))
	ErrLookingUpPluginSymbol = NewErrorWrapper("ErrLookingUpPluginSymbol", errors.New("error looking up plugin symbol from file %s"))

	// General errors
	ErrHashingPassword       = NewErrorWrapper("ErrHashingPassword", errors.New("error hashing password"))
	ErrInvalidRequestPayload = NewErrorWrapper("ErrInvalidRequestPayload", errors.New("invalid request payload"))
	ErrOperationNotPermitted = NewErrorWrapper("ErrOperationNotPermitted", errors.New("operation not permitted"))

	// Uncategorized errors
	ErrInvalidHookConfig = NewErrorWrapper("ErrInvalidHookConfig", errors.New("invalid hook config"))
	ErrTest              = NewErrorWrapper("ErrTest", errors.New("error test %s, %d, %f, %t"))
)
