package httphandler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

var ErrInvalidQueryParam = errors.New("query param is invalid")

type streetMarketLister interface {
	List(context.Context, int, domain.StreetMarketFilter) ([]domain.StreetMarket, *domain.Error)
}

type streetMarketListHandlerLogger interface {
	Error(context.Context, domain.Error)
}

type listStreetMarketResponse map[string][]streetMarketResponse

type StreetMarketListHandler struct {
	getter streetMarketLister
	logger streetMarketListHandlerLogger
}

func NewStreetMarketListHandler(getter streetMarketLister, logger streetMarketListHandlerLogger) *StreetMarketListHandler {
	return &StreetMarketListHandler{getter, logger}
}

func (h *StreetMarketListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	f := domain.StreetMarketFilter{
		District:     r.FormValue("district"),
		Region5:      r.FormValue("region5"),
		Name:         r.FormValue("name"),
		Neighborhood: r.FormValue("neighborhood"),
	}

	var pgn int
	var err error
	page := r.FormValue("page")
	if page != "" {
		pgn, err = strconv.Atoi(page)
		if err != nil {
			h.logger.Error(ctx, domain.Error{
				Kind: domain.InpValidationErrKd,
				Msg:  err.Error(),
			})
			respondError(w, http.StatusBadRequest, "Page can be integer")
			return
		}
	}

	ls, dErr := h.getter.List(ctx, pgn, f)
	if dErr != nil {
		h.logger.Error(ctx, *dErr)
		respondError(w, http.StatusInternalServerError, dErr.Error())
		return
	}

	lr := []streetMarketResponse{}
	for _, sm := range ls {
		smr := streetMarketResponse{
			ID:            sm.ID,
			Long:          sm.Long,
			Lat:           sm.Lat,
			SectCens:      sm.SectCens,
			Area:          sm.Area,
			IDdist:        sm.IDdist,
			District:      sm.District,
			IDSubTH:       sm.IDSubTH,
			SubTownHall:   sm.SubTownHall,
			Region5:       sm.Region5,
			Region8:       sm.Region8,
			Name:          sm.Name,
			Register:      sm.Register,
			Street:        sm.Street,
			Number:        sm.Number,
			Neighborhood:  sm.Neighborhood,
			AddrExtraInfo: sm.AddrExtraInfo,
		}

		lr = append(lr, smr)
	}

	respondJSON(w, http.StatusOK, listStreetMarketResponse{"data": lr})
}
