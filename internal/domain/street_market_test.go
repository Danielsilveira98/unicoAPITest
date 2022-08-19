package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestStreetMarketGetInput_Validate(t *testing.T) {
	id := uuid.NewString()
	smgi := StreetMarketGetInput{ID: id}
	got := smgi.Validate()

	if got != nil {
		t.Errorf("Want nil, got %v", got)
	}
}

func TestStreetMarketGetInput_Validate_Error(t *testing.T) {
	id := "Invalid"
	smgi := StreetMarketGetInput{ID: id}
	got := smgi.Validate()

	if !errors.Is(got, ErrInpValidation) {
		t.Errorf("Want error %v, got error %v", ErrInpValidation, got)
	}
}
