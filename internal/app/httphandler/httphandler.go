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
	SectCens      string  `json:"sect_cens"`
	Area          string  `json:"area"`
	IDdist        string  `json:"id_dist"`
	District      string  `json:"district"`
	IDSubTH       string  `json:"id_sub_th"`
	SubTownHall   string  `json:"subtownhall"`
	Region5       string  `json:"region_5"`
	Region8       string  `json:"region_8"`
	Name          string  `json:"name"`
	Register      string  `json:"register"`
	Street        string  `json:"street"`
	Number        string  `json:"number"`
	Neighborhood  string  `json:"neighborhood"`
	AddrExtraInfo string  `json:"addr_extra_info"`
}

type streetMarketResponse struct {
	ID            string  `json:"id,omitempty"`
	Long          float64 `json:"long,omitempty"`
	Lat           float64 `json:"lat,omitempty"`
	SectCens      string  `json:"sect_cens,omitempty"`
	Area          string  `json:"area,omitempty"`
	IDdist        string  `json:"id_dist,omitempty"`
	District      string  `json:"district,omitempty"`
	IDSubTH       string  `json:"id_sub_th,omitempty"`
	SubTownHall   string  `json:"subtownhall,omitempty"`
	Region5       string  `json:"region_5,omitempty"`
	Region8       string  `json:"region_8,omitempty"`
	Name          string  `json:"name,omitempty"`
	Register      string  `json:"register,omitempty"`
	Street        string  `json:"street,omitempty"`
	Number        string  `json:"number,omitempty"`
	Neighborhood  string  `json:"neighborhood,omitempty"`
	AddrExtraInfo string  `json:"addr_extra_info,omitempty"`
}
