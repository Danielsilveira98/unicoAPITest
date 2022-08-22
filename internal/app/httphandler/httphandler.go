package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse map[string]interface{}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println(err) // TODO log here
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(response)); err != nil {
		fmt.Println(err) // TODO log here
	}
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, ErrorResponse{"error": message})
}

type streetMarketBody struct {
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
