package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type StreetMarketGetQuery struct {
	District     string
	Region5      string
	Name         string
	Neighborhood string
}

type StreetMarketGetInput struct {
	ID    string
	Query StreetMarketGetQuery
}

func (i *StreetMarketGetInput) Validate() error {
	if _, err := uuid.Parse(i.ID); err != nil {
		return fmt.Errorf("%w The ID can be a valid UUID.", ErrInpValidation)
	}

	return nil
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
