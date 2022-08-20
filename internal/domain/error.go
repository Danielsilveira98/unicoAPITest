package domain

import "errors"

var (
	ErrInpValidation       = errors.New("Input is invalid.")
	ErrUnexpected          = errors.New("Unexpected error.")
	ErrOrgCalNotFound      = errors.New("Organization calendar not found.")
	ErrForeignKeyNotExists = errors.New("Foreign key not exists.")
	ErrNothingUpdated      = errors.New("Nothing updated.")
	ErrSMNotFound          = errors.New("Street Market not found.")
	ErrNothingDeleted      = errors.New("Nothing deleted.")
)
