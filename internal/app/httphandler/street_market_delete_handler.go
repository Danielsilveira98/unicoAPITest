package httphandler

import (
	"context"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/gorilla/mux"
)

type streetMarketEraser interface {
	Delete(context.Context, domain.SMID) *domain.Error
}

type streetMarketDeleteHandlerLogger interface {
	Error(context.Context, domain.Error)
}

type StreetMarketDeleteHandler struct {
	eraser streetMarketEraser
	logger streetMarketWriteHandlerLogger
}

func NewStreetMarketDeleteHandler(
	repo streetMarketEraser,
	logger streetMarketWriteHandlerLogger,
) *StreetMarketDeleteHandler {
	return &StreetMarketDeleteHandler{repo, logger}
}

func (h *StreetMarketDeleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := domain.SMID(vars["street-market-id"])

	err := h.eraser.Delete(r.Context(), id)

	if err != nil {
		var status int

		switch err.Kind {
		case domain.InpValidationErrKd:
			status = http.StatusBadRequest
		case domain.SMNotFoundErrKd:
			status = http.StatusNotFound
		default:
			h.logger.Error(ctx, *err)
			status = http.StatusInternalServerError
		}

		respondError(w, status, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, "")
}
