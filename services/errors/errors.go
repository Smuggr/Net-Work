package errors

import "errors"

var (
	ErrInsufficientLength    = errors.New("insufficient length")
	ErrInsufficientUppercase = errors.New("insufficient uppercase letters")
	ErrInsufficientLowercase = errors.New("insufficient lowercase letters")
	ErrInsufficientDigits    = errors.New("insufficient digits")
	ErrInsufficientSpecial   = errors.New("insufficient special characters")
)
