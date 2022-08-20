package domain

import (
	"fmt"
)

type StreetMarketFilter struct {
	District     string
	Region5      string
	Name         string
	Neighborhood string
}

type StreetMarket struct {
	ID            string  // 	ID
	Long          float64 // LONG
	Lat           float64 // LAT
	SectCens      string  // SETCENS
	Area          string  // AREAP
	IDdist        int     // CODDIST
	District      string  // DISTRITO
	IDSubTH       int     // CODSUBPREF
	SubTownHall   string  // SUBPREF
	Region5       string  // REGIAO5
	Region8       string  // REGIAO8
	Name          string  // NOME_FEIRA
	Register      string  // REGISTRO
	Street        string  // LOGRADOURO
	Number        int     // NUMERO
	Neighborhood  string  // BAIRRO
	AddrExtraInfo string  // REFERENCIA
}

type StreetMarketCreateInput struct {
	Long          float64 // LONG
	Lat           float64 // LAT
	SectCens      string  // SETCENS
	Area          string  // AREAP
	IDdist        int     // CODDIST
	District      string  // DISTRITO
	IDSubTH       int     // CODSUBPREF
	SubTownHall   string  // SUBPREF
	Region5       string  // REGIAO5
	Region8       string  // REGIAO8
	Name          string  // NOME_FEIRA
	Register      string  // REGISTRO
	Street        string  // LOGRADOURO
	Number        int     // NUMERO
	Neighborhood  string  // BAIRRO
	AddrExtraInfo string  // REFERENCIA
}

func (i *StreetMarketCreateInput) Validate() error {
	if i.Name == "" {
		return fmt.Errorf("%w Name is required", ErrInpValidation)
	}

	if i.Register == "" {
		return fmt.Errorf("%w Register is required", ErrInpValidation)
	}

	if i.Street == "" {
		return fmt.Errorf("%w Street is required", ErrInpValidation)
	}

	if i.Neighborhood == "" {
		return fmt.Errorf("%w Neighborhood is required", ErrInpValidation)
	}

	return nil
}

type StreetMarketEditInput struct {
	Long          float64 // LONG
	Lat           float64 // LAT
	SectCens      string  // SETCENS
	Area          string  // AREAP
	IDdist        int     // CODDIST
	District      string  // DISTRITO
	IDSubTH       int     // CODSUBPREF
	SubTownHall   string  // SUBPREF
	Region5       string  // REGIAO5
	Region8       string  // REGIAO8
	Name          string  // NOME_FEIRA
	Register      string  // REGISTRO
	Street        string  // LOGRADOURO
	Number        int     // NUMERO
	Neighborhood  string  // BAIRRO
	AddrExtraInfo string  // REFERENCIA
}
