package domain

import (
	"fmt"

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
	IDdist        int
	District      string
	IDSubTH       int
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

type StreetMarketCreateInput struct {
	Long          float64
	Lat           float64
	SectCens      string
	Area          string
	IDdist        int
	District      string
	IDSubTH       int
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

func (i *StreetMarketCreateInput) Validate() error {
	return nil
}

type StreetMarketEditInput struct {
	Long          float64
	Lat           float64
	SectCens      string
	Area          string
	IDdist        int
	District      string
	IDSubTH       int
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
