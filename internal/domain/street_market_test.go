package domain

import (
	"errors"
	"testing"
)

func TestSMID_Validate(t *testing.T) {
	var smid SMID = "7f236270-6c0d-47bb-9e8c-86315865f8a0"

	err := smid.Validate()

	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}
}

func TestSMID_Validate_Error(t *testing.T) {
	var smid SMID = "invalid"

	err := smid.Validate()

	if !errors.Is(err, ErrInpValidation) {
		t.Errorf("want error %v, got %v", ErrInpValidation, err)
	}
}
