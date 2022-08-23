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
	create := StreetMarketCreateInput{
		Long:          -46548146,
		Lat:           -23568390,
		SectCens:      "355030885000019",
		Area:          "3550308005040",
		IDdist:        "87",
		District:      "VILA FORMOSA",
		IDSubTH:       "26",
		SubTownHall:   "ARICANDUVA",
		Region5:       "Leste",
		Region8:       "Leste 1",
		Name:          "RAPOSO TAVARES",
		Register:      "1129-0",
		Street:        "Rua dos Bobos",
		Number:        "500",
		Neighborhood:  "JARDIM SARAH",
		AddrExtraInfo: "Loren ipsum",
	}

	if err := create.Validate(); err != nil {
		t.Errorf("expect nil, got %v", err)
	}
}

func TestStreetMarketCreateInput_Validate_Error(t *testing.T) {
	testCases := map[string]StreetMarketCreateInput{
		"When is empty": {},
		"When Lat is empty": {
			Long: -46548146},
		"When SectCens is empty": {
			Long: -46548146,
			Lat:  -23568390},
		"When Area is empty": {
			Long:     -46548146,
			Lat:      -23568390,
			SectCens: "355030885000019"},
		"When IDdist is empty": {
			Long:     -46548146,
			Lat:      -23568390,
			SectCens: "355030885000019",
			Area:     "3550308005040"},
		"When District is empty": {
			Long:     -46548146,
			Lat:      -23568390,
			SectCens: "355030885000019",
			Area:     "3550308005040",
			IDdist:   "87"},
		"When IDSubTH is empty": {
			Long:     -46548146,
			Lat:      -23568390,
			SectCens: "355030885000019",
			Area:     "3550308005040",
			IDdist:   "87",
			District: "VILA FORMOSA",
		},
		"When SubTownHall is empty": {
			Long:     -46548146,
			Lat:      -23568390,
			SectCens: "355030885000019",
			Area:     "3550308005040",
			IDdist:   "87",
			District: "VILA FORMOSA",
			IDSubTH:  "26",
		},
		"When Region5 is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
		},
		"When Region8 is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
		},
		"When Name is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
			Region8:     "Leste 1",
		},
		"When Register is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
			Region8:     "Leste 1",
			Name:        "RAPOSO TAVARES",
		},
		"When Street is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
			Region8:     "Leste 1",
			Name:        "RAPOSO TAVARES",
			Register:    "1129-0",
		},
		"When Number is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
			Region8:     "Leste 1",
			Name:        "RAPOSO TAVARES",
			Register:    "1129-0",
			Street:      "Rua dos Bobos",
		},
		"When Neighborhood is empty": {
			Long:        -46548146,
			Lat:         -23568390,
			SectCens:    "355030885000019",
			Area:        "3550308005040",
			IDdist:      "87",
			District:    "VILA FORMOSA",
			IDSubTH:     "26",
			SubTownHall: "ARICANDUVA",
			Region5:     "Leste",
			Region8:     "Leste 1",
			Name:        "RAPOSO TAVARES",
			Register:    "1129-0",
			Street:      "Rua dos Bobos",
			Number:      "500",
		},
		"When AddrExtraInfo is empty": {
			Long:         -46548146,
			Lat:          -23568390,
			SectCens:     "355030885000019",
			Area:         "3550308005040",
			IDdist:       "87",
			District:     "VILA FORMOSA",
			IDSubTH:      "26",
			SubTownHall:  "ARICANDUVA",
			Region5:      "Leste",
			Region8:      "Leste 1",
			Name:         "RAPOSO TAVARES",
			Register:     "1129-0",
			Street:       "Rua dos Bobos",
			Number:       "500",
			Neighborhood: "JARDIM SARAH",
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			err := tc.Validate()
			if err == nil {
				t.Error("expect err, got nil")
			}
		})
	}
}
