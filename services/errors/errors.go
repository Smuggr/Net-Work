package errors

import "errors"

var (
	ErrInsufficientLength    = errors.New("insufficient length")
	ErrInsufficientUppercase = errors.New("insufficient uppercase letters")
	ErrInsufficientLowercase = errors.New("insufficient lowercase letters")
	ErrInsufficientDigits    = errors.New("insufficient digits")
	ErrInsufficientSpecial   = errors.New("insufficient special characters")

	ErrHashingPassword       = errors.New("error hashing password")
	ErrInvalidRequestPayload = errors.New("invalid request payload")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrCreatingToken         = errors.New("error creating token")

	ErrDefaultAdminNotFound  = errors.New("default admin not found")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUserNotFound          = errors.New("user not found")
)
