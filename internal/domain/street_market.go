package domain

import (
	"time"

	"github.com/google/uuid"
)

type SMID string

func (s *SMID) Validate() *Error {
	if _, err := uuid.Parse(string(*s)); err != nil {
		return &Error{Kind: InpValidationErrKd, Msg: err.Error()}
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

func (d *StreetMarketCreateInput) Validate() *Error {
	if d.Long == 0.0 {
		return &Error{Kind: InpValidationErrKd, Msg: "Long is required"}
	}
	if d.Lat == 0.0 {
		return &Error{Kind: InpValidationErrKd, Msg: "Lat is required"}
	}
	if d.SectCens == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "SectCens is required"}
	}
	if d.Area == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Area is required"}
	}
	if d.IDdist == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "IDdist is required"}
	}
	if d.District == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "District is required"}
	}
	if d.IDSubTH == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "IDSubTH is required"}
	}
	if d.SubTownHall == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "SubTownHall is required"}
	}
	if d.Region5 == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Region5 is required"}
	}
	if d.Region8 == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Region8 is required"}
	}
	if d.Name == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Name is required"}
	}
	if d.Register == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Register is required"}
	}
	if d.Street == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Street is required"}
	}
	if d.Number == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Number is required"}
	}
	if d.Neighborhood == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "Neighborhood is required"}
	}
	if d.AddrExtraInfo == "" {
		return &Error{Kind: InpValidationErrKd, Msg: "AddrExtraInfo is required"}
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
