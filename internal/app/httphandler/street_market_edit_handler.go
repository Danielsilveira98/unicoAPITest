package httphandler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/gorilla/mux"
)

type streetMarketEditor interface {
	Edit(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) *domain.Error
}

type streetMarketWriteHandlerLogger interface {
	Error(context.Context, domain.Error)
}

type StreetMarketEditHandler struct {
	editor streetMarketEditor
	logger streetMarketWriteHandlerLogger
}

func NewStreetMarketEditHandler(editor streetMarketEditor, logger streetMarketWriteHandlerLogger) *StreetMarketEditHandler {
	return &StreetMarketEditHandler{editor, logger}
}

func (h *StreetMarketEditHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

	dErr := h.editor.Edit(r.Context(), id, input)

	if dErr != nil {
		var status int

		switch dErr.Kind {
		case domain.InpValidationErrKd:
			status = http.StatusBadRequest
		case domain.SMNotFoundErrKd:
			status = http.StatusNotFound
		default:
			h.logger.Error(ctx, *dErr)
			status = http.StatusInternalServerError
		}

		respondError(w, status, dErr.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, "")
}
