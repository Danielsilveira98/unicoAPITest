package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/gorilla/mux"
)

type streetMarketEditor interface {
	Edit(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) error
}

type StreetMarketEditHandler struct {
	editor streetMarketEditor
}

func NewStreetMarketEditHandler(editor streetMarketEditor) *StreetMarketEditHandler {
	return &StreetMarketEditHandler{editor}
}

func (h *StreetMarketEditHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body streetMarketBody

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

	vars := mux.Vars(r)
	id := domain.SMID(vars["street-market-id"])

	input := domain.StreetMarketEditInput{
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

	err = h.editor.Edit(r.Context(), id, input)

	if err != nil {
		fmt.Printf("%v", err)
		var status int

		switch err {
		case domain.ErrInpValidation:
			status = http.StatusBadRequest
		case domain.ErrSMNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}

		respondError(w, status, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, "")
}
