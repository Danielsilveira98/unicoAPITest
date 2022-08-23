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
	List(context.Context, int, domain.StreetMarketFilter) ([]domain.StreetMarket, error)
}

type listStreetMarketResponse map[string][]streetMarketResponse

type StreetMarketListHandler struct {
	getter streetMarketLister
}

func NewStreetMarketListHandler(getter streetMarketLister) *StreetMarketListHandler {
	return &StreetMarketListHandler{getter}
}

func (h *StreetMarketListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	f := domain.StreetMarketFilter{
		District:     r.FormValue("distrito"),
		Region5:      r.FormValue("regiao5"),
		Name:         r.FormValue("nome_feira"),
		Neighborhood: r.FormValue("bairro"),
	}

	var pgn int
	var err error
	page := r.FormValue("page")
	if page != "" {
		pgn, err = strconv.Atoi(page)
		if err != nil {
			respondError(w, http.StatusBadRequest, ErrInvalidQueryParam.Error())
			return
		}
	}

	ls, err := h.getter.List(r.Context(), pgn, f)

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
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
