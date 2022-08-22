package httphandler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
)

type stubStreetMarketCreator struct {
	createInp domain.StreetMarketCreateInput
	create    func(context.Context, domain.StreetMarketCreateInput) (string, error)
}

func (s *stubStreetMarketCreator) Create(ctx context.Context, inp domain.StreetMarketCreateInput) (string, error) {
	s.createInp = inp
	return s.create(ctx, inp)
}

func TestStreetMarketCreateHandler_Handle(t *testing.T) {
	id := "12d54a54-bbd7-4e70-8c3c-7e8e424d8ebf"
	wantInp := domain.StreetMarketCreateInput{
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
	}
	requestBody := createSMBody{
		wantInp.Long,
		wantInp.Lat,
		wantInp.SectCens,
		wantInp.Area,
		wantInp.IDdist,
		wantInp.District,
		wantInp.IDSubTH,
		wantInp.SubTownHall,
		wantInp.Region5,
		wantInp.Region8,
		wantInp.Name,
		wantInp.Register,
		wantInp.Street,
		wantInp.Number,
		wantInp.Neighborhood,
		wantInp.AddrExtraInfo,
	}

	creatorMock := &stubStreetMarketCreator{
		create: func(ctx context.Context, pci domain.StreetMarketCreateInput) (string, error) {
			return id, nil
		},
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(requestBody); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/street_market", &body)
	if err != nil {
		t.Fatal(err)
	}

	h := NewStreetMarketCreateHandler(creatorMock)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expect status code %v, got %v", http.StatusCreated, status)
	}

	if diff := cmp.Diff(wantInp, creatorMock.createInp); diff != "" {
		t.Errorf("street market creator create receive a unexpected input (-want +got):\n%s", diff)
	}
}

func TestStreetMarketCreateHandler_Handle_Error(t *testing.T) {
	testCases := map[string]struct {
		rBody        createSMBody
		creatorErr   error
		wantStatusCd int
		wantBody     ErrorResponse
	}{
		"Unexpected error": {
			rBody:        createSMBody{},
			creatorErr:   domain.ErrUnexpected,
			wantStatusCd: http.StatusInternalServerError,
			wantBody:     ErrorResponse{"error": domain.ErrUnexpected.Error()},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			creatorMock := &stubStreetMarketCreator{
				create: func(ctx context.Context, pci domain.StreetMarketCreateInput) (string, error) {
					return "", tc.creatorErr
				},
			}

			var body bytes.Buffer
			if err := json.NewEncoder(&body).Encode(tc.rBody); err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("POST", "/street_market", &body)
			if err != nil {
				t.Fatal(err)
			}

			h := NewStreetMarketCreateHandler(creatorMock)
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

	t.Run("malformed body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/street_market", strings.NewReader("body"))
		if err != nil {
			t.Fatal(err)
		}

		h := NewStreetMarketCreateHandler(&stubStreetMarketCreator{})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.Handle)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("expect status code %v, got %v", status, http.StatusBadRequest)
		}
		var got ErrorResponse
		err = json.Unmarshal(rr.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}

		want := ErrorResponse{"error": "malformed body"}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("want body mismatch with got body (-want +got):\n%s", diff)
		}
	})
}
