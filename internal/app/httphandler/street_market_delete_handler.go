package httphandler

import (
	"context"
	"net/http"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/gorilla/mux"
)

type streetMarketEraser interface {
	Delete(context.Context, domain.SMID) error
}

type StreetMarketDeleteHandler struct {
	eraser streetMarketEraser
}

func NewStreetMarketDeleteHandler(repo streetMarketEraser) *StreetMarketDeleteHandler {
	return &StreetMarketDeleteHandler{repo}
}

func (h *StreetMarketDeleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := domain.SMID(vars["street-market-id"])

	err := h.eraser.Delete(r.Context(), id)

	if err != nil {
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
