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

func TestStreetMarketCreateInput_Validate(t *testing.T) {
	smin := StreetMarketCreateInput{
		Name:         "Joh",
		Register:     "22222",
		Street:       "Street",
		Neighborhood: "Neighborhood",
	}

	if err := smin.Validate(); err != nil {
		t.Errorf("expect return nil, got %v", err)
	}
}
func TestStreetMartketCreateInput_Validate_Error(t *testing.T) {
	testCases := map[string]struct {
		wErr error
		inp  StreetMarketCreateInput
	}{
		"When name is empty": {
			wErr: ErrInpValidation,
			inp:  StreetMarketCreateInput{},
		},
		"When Register is empty": {
			wErr: ErrInpValidation,
			inp: StreetMarketCreateInput{
				Name: "Joh",
			},
		},
		"When Street is empty": {
			wErr: ErrInpValidation,
			inp: StreetMarketCreateInput{
				Name:     "Joh",
				Register: "22222",
			},
		},
		"When Neighborhood is empty": {
			wErr: ErrInpValidation,
			inp: StreetMarketCreateInput{
				Name:     "Joh",
				Register: "22222",
				Street:   "Ali",
			},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			if gErr := tc.inp.Validate(); !errors.Is(gErr, tc.wErr) {
				t.Errorf("Want error %v, got error %v", tc.wErr, gErr)
			}
		})
	}
}
