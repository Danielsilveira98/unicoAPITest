package domain

import "errors"

var (
	ErrInpValidation = errors.New("Input is invalid.")
	ErrUnexpected    = errors.New("Unexpected error.")
)
