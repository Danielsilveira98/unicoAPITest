package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type createSMBody struct {
	Long          float64 `json:"long"`
	Lat           float64 `json:"lat"`
	SectCens      string  `json:"setcens"`
	Area          string  `json:"areap"`
	IDdist        int     `json:"coddist"`
	District      string  `json:"distrito"`
	IDSubTH       int     `json:"codsubpref"`
	SubTownHall   string  `json:"subpref"`
	Region5       string  `json:"regiao5"`
	Region8       string  `json:"regiao8"`
	Name          string  `json:"nome_feira"`
	Register      string  `json:"registro"`
	Street        string  `json:"logradouro"`
	Number        string  `json:"numero"`
	Neighborhood  string  `json:"bairro"`
	AddrExtraInfo string  `json:"referencia"`
}

type streetMarketCreator interface {
	Create(context.Context,
		domain.StreetMarketCreateInput) (string,
		error)
}

type StreetMarketCreateHandler struct {
	creator streetMarketCreator
}

func NewStreetMarketCreateHandler(creator streetMarketCreator) *StreetMarketCreateHandler {
	return &StreetMarketCreateHandler{creator}
}

func (h *StreetMarketCreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body createSMBody

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bb, &body); err != nil {
		respondError(w, http.StatusBadRequest, "malformed body")
		return
	}

	input := domain.StreetMarketCreateInput{
		Long:          body.Long,
		Lat:           body.Lat,
		SectCens:      body.SectCens,
		Area:          body.Area,
		IDdist:        body.IDdist,
		District:      body.District,
		IDSubTH:       body.IDSubTH,
		SubTownHall:   body.SubTownHall,
		Region5:       body.Region5,
		Region8:       body.Region8,
		Name:          body.Name,
		Register:      body.Register,
		Street:        body.Street,
		Number:        body.Number,
		Neighborhood:  body.Neighborhood,
		AddrExtraInfo: body.AddrExtraInfo,
	}

	id, err := h.creator.Create(r.Context(), input)

	if err != nil {
		fmt.Printf("%v", err)
		var status int

		switch err {
		case domain.ErrInpValidation:
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}

		respondError(w, status, err.Error())
		return
	}

	location := fmt.Sprintf("%s/street_market/%s", r.Host, id)
	w.Header().Add("Location", location)
	respondJSON(w, http.StatusCreated, "")
}
