package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

type stubStreetMarketLister struct {
	listInp domain.StreetMarketFilter
	list    func(context.Context, domain.StreetMarketFilter) ([]domain.StreetMarket, error)
}

func (s *stubStreetMarketLister) List(
	ctx context.Context,
	inp domain.StreetMarketFilter,
) ([]domain.StreetMarket, error) {
	s.listInp = inp
	return s.list(ctx, inp)
}

func TestStreetMarketListHandler_Handle(t *testing.T) {
	list := []domain.StreetMarket{{
		ID:            "2c809e53-6e2e-4a60-bbf4-de8913562970",
		Long:          -46548146,
		Lat:           -23568390,
		SectCens:      "355030885000019",
		Area:          "3550308005040",
		IDdist:        87,
		District:      "VILA FORMOSA",
		IDSubTH:       26,
		SubTownHall:   "ARICANDUVA",
		Region5:       "Leste",
		Region8:       "Leste 1",
		Name:          "RAPOSO TAVARES",
		Register:      "1129-0",
		Street:        "Rua dos Bobos",
		Number:        "500",
		Neighborhood:  "JARDIM SARAH",
		AddrExtraInfo: "Loren ipsum",
	}}

	want := listStreetMarketResponse{
		"data": []streetMarketResponse{{
			ID:            list[0].ID,
			Long:          list[0].Long,
			Lat:           list[0].Lat,
			SectCens:      list[0].SectCens,
			Area:          list[0].Area,
			IDdist:        list[0].IDdist,
			District:      list[0].District,
			IDSubTH:       list[0].IDSubTH,
			SubTownHall:   list[0].SubTownHall,
			Region5:       list[0].Region5,
			Region8:       list[0].Region8,
			Name:          list[0].Name,
			Register:      list[0].Register,
			Street:        list[0].Street,
			Number:        list[0].Number,
			Neighborhood:  list[0].Neighborhood,
			AddrExtraInfo: list[0].AddrExtraInfo,
		}},
	}

	listerMock := &stubStreetMarketLister{
		list: func(ctx context.Context, inp domain.StreetMarketFilter) ([]domain.StreetMarket, error) {
			return list, nil
		},
	}

	wantInp := domain.StreetMarketFilter{
		District:     "distrito",
		Neighborhood: "centro",
	}

	path := fmt.Sprintf("/street_market?distrito=%s&bairro=%s", wantInp.District, wantInp.Neighborhood)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	h := NewStreetMarketHandler(listerMock)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expect status code %v, got %v", status, http.StatusOK)
	}

	var got listStreetMarketResponse
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("want body mismatch with got body (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(wantInp, listerMock.listInp); diff != "" {
		t.Errorf("street market lister list receive a unexpected input  (-want +got):\n%s", diff)
	}
}

func TestStreetMarketListHandler_Handle_Error(t *testing.T) {
	testCases := map[string]struct {
		listerErr    error
		wantStatusCd int
		wantBody     ErrorResponse
	}{
		"Unexpected error": {
			listerErr:    domain.ErrUnexpected,
			wantStatusCd: http.StatusInternalServerError,
			wantBody:     ErrorResponse{"error": domain.ErrUnexpected.Error()},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			listerMock := &stubStreetMarketLister{
				list: func(ctx context.Context, inp domain.StreetMarketFilter) ([]domain.StreetMarket, error) {
					return nil, tc.listerErr
				},
			}

			req, err := http.NewRequest(http.MethodGet, "/street_market", nil)
			if err != nil {
				t.Fatal(err)
			}

			h := NewStreetMarketHandler(listerMock)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.Handle)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatusCd {
				t.Errorf("expect status code %v, got %v", tc.wantStatusCd, status)
			}

			var got ErrorResponse
			err = json.Unmarshal(rr.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.wantBody, got); diff != "" {
				t.Errorf("want body mismatch with got body (-want +got):\n%s", diff)
			}
		})
	}
}
