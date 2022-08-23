package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SMID string

func (s *SMID) Validate() error {
	if _, err := uuid.Parse(string(*s)); err != nil {
		return fmt.Errorf("%w IS is an invalid UUID.", ErrInpValidation)
	}

	return nil
}

type StreetMarketFilter struct {
	District     string
	Region5      string
	Name         string
	Neighborhood string
}

type StreetMarket struct {
	ID            string
	Long          float64
	Lat           float64
	SectCens      string
	Area          string
	IDdist        string
	District      string
	IDSubTH       string
	SubTownHall   string
	Region5       string
	Region8       string
	Name          string
	Register      string
	Street        string
	Number        string
	Neighborhood  string
	AddrExtraInfo string
	CreatedAt     *time.Time
}

type StreetMarketCreateInput struct {
	Long          float64
	Lat           float64
	SectCens      string
	Area          string
	IDdist        string
	District      string
	IDSubTH       string
	SubTownHall   string
	Region5       string
	Region8       string
	Name          string
	Register      string
	Street        string
	Number        string
	Neighborhood  string
	AddrExtraInfo string
}

func (d *StreetMarketCreateInput) Validate() error {
	if d.Long == 0.0 {
		return fmt.Errorf("%w Long is required", ErrInpValidation)
	}
	if d.Lat == 0.0 {
		return fmt.Errorf("%w Lat is required", ErrInpValidation)
	}
	if d.SectCens == "" {
		return fmt.Errorf("%w SectCens is required", ErrInpValidation)
	}
	if d.Area == "" {
		return fmt.Errorf("%w Area is required", ErrInpValidation)
	}
	if d.IDdist == "" {
		return fmt.Errorf("%w IDdist is required", ErrInpValidation)
	}
	if d.District == "" {
		return fmt.Errorf("%w District is required", ErrInpValidation)
	}
	if d.IDSubTH == "" {
		return fmt.Errorf("%w IDSubTH is required", ErrInpValidation)
	}
	if d.SubTownHall == "" {
		return fmt.Errorf("%w SubTownHall is required", ErrInpValidation)
	}
	if d.Region5 == "" {
		return fmt.Errorf("%w Region5 is required", ErrInpValidation)
	}
	if d.Region8 == "" {
		return fmt.Errorf("%w Region8 is required", ErrInpValidation)
	}
	if d.Name == "" {
		return fmt.Errorf("%w Name is required", ErrInpValidation)
	}
	if d.Register == "" {
		return fmt.Errorf("%w Register is required", ErrInpValidation)
	}
	if d.Street == "" {
		return fmt.Errorf("%w Street is required", ErrInpValidation)
	}
	if d.Number == "" {
		return fmt.Errorf("%w Number is required", ErrInpValidation)
	}
	if d.Neighborhood == "" {
		return fmt.Errorf("%w Neighborhood is required", ErrInpValidation)
	}
	if d.AddrExtraInfo == "" {
		return fmt.Errorf("%w AddrExtraInfo is required", ErrInpValidation)
	}

	return nil
}

type StreetMarketEditInput struct {
	Long          float64
	Lat           float64
	SectCens      string
	Area          string
	IDdist        string
	District      string
	IDSubTH       string
	SubTownHall   string
	Region5       string
	Region8       string
	Name          string
	Register      string
	Street        string
	Number        string
	Neighborhood  string
	AddrExtraInfo string
}

type Pagination struct {
	Offset int
	Limit  int
}
