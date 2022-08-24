package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type streetMarketCreator interface {
	Create(context.Context,
		domain.StreetMarketCreateInput) (string, *domain.Error)
}
type streetMarketCreateHandlerLogger interface {
	Error(context.Context, domain.Error)
}

type StreetMarketCreateHandler struct {
	creator streetMarketCreator
	logger  streetMarketCreateHandlerLogger
}

func NewStreetMarketCreateHandler(
	creator streetMarketCreator,
	logger streetMarketCreateHandlerLogger,
) *StreetMarketCreateHandler {
	return &StreetMarketCreateHandler{creator, logger}
}

func (h *StreetMarketCreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body streetMarketBody

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(ctx, domain.Error{
			Kind: domain.UnexpectedErrKd,
			Msg:  err.Error(),
		})
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bb, &body); err != nil {
		h.logger.Error(ctx, domain.Error{
			Kind: domain.InpValidationErrKd,
			Msg:  err.Error(),
		})
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

	id, dErr := h.creator.Create(r.Context(), input)

	if dErr != nil {
		var status int

		switch dErr.Kind {
		case domain.InpValidationErrKd:
			status = http.StatusBadRequest
		default:
			h.logger.Error(ctx, *dErr)
			status = http.StatusInternalServerError
		}

		respondError(w, status, dErr.Error())
		return
	}

	location := fmt.Sprintf("%s/street_market/%s", r.Host, id)
	w.Header().Add("Location", location)
	respondJSON(w, http.StatusCreated, "")
}
