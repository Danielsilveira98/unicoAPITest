package domain

import "errors"

type KindError string

const (
	UnexpectedErrKd     KindError = "UNEXPECTED"
	NothingCreatedErrKd KindError = "NOTHING_CREATED"
	NothingUpdatedErrKd KindError = "NOTHING_UPDATED"
	NothingDeletedErrKd KindError = "NOTHING_DELETED"
	SMNotFoundErrKd     KindError = "STREET_MARKET_NOT_FOUND"
	InpValidationErrKd  KindError = "INPUT_IS_INVALID"
)

type Error struct {
	Kind     KindError
	Msg      string
	Previous *Error
}

func (e *Error) Error() string {
	return e.Msg
}

var (
	ErrInpValidation  = errors.New("Input is invalid.")
	ErrUnexpected     = errors.New("Unexpected error.")
	ErrSMNotFound     = errors.New("Street Market not found.")
	ErrNothingUpdated = errors.New("Nothing updated.")
	ErrNothingDeleted = errors.New("Nothing deleted.")
	ErrNothingCreated = errors.New("Nothing created.")
)
