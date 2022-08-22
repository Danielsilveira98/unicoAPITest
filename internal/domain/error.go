package domain

import "errors"

var (
	ErrInpValidation  = errors.New("Input is invalid.")
	ErrUnexpected     = errors.New("Unexpected error.")
	ErrSMNotFound     = errors.New("Street Market not found.")
	ErrNothingUpdated = errors.New("Nothing updated.")
	ErrNothingDeleted = errors.New("Nothing deleted.")
	ErrNothingCreated = errors.New("Nothing created.")
)
