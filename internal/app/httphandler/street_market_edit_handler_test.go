package httphandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type stubStreetMarketEditor struct {
	editIDInp domain.SMID
	editInp   domain.StreetMarketEditInput
	edit      func(context.Context, domain.SMID, domain.StreetMarketEditInput) error
}

func (s *stubStreetMarketEditor) Edit(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) error {
	s.editInp = inp
	s.editIDInp = ID
	return s.edit(ctx, ID, inp)
}

func TestStreetMarketEditHandler_Handle(t *testing.T) {
	id := "aaa1be24-ddec-4590-839d-b7ae54b9ed78"
	wantInp := domain.StreetMarketEditInput{
		Long:          -46548146,
		Lat:           -23568390,
		SectCens:      "355030885000019",
		Area:          "3550308005040",
		IDdist:        "87",
		District:      "VILA FORMOSA",
		IDSubTH:       "26",
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
	requestBody := streetMarketBody{
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

	editorMock := &stubStreetMarketEditor{
		edit: func(ctx context.Context, ID domain.SMID, inp domain.StreetMarketEditInput) error {
			return nil
		},
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(requestBody); err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("/street_market/%s", id)
	req, err := http.NewRequest(http.MethodPatch, path, &body)
	if err != nil {
		t.Fatal(err)
	}

	h := NewStreetMarketEditHandler(editorMock)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/street_market/{street-market-id}", h.Handle)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expect status code %v, got %v", http.StatusNoContent, status)
	}

	if diff := cmp.Diff(wantInp, editorMock.editInp); diff != "" {
		t.Errorf("street market editor edit receive a unexpected input (-want +got):\n%s", diff)
	}
}

func TestStreetMarketEditHandler_Handle_Error(t *testing.T) {
	testCases := map[string]struct {
		rBody        streetMarketBody
		id           string
		editorErr    error
		wantStatusCd int
		wantBody     ErrorResponse
	}{
		"Invalid id": {
			rBody:        streetMarketBody{},
			id:           "invalid",
			editorErr:    domain.ErrInpValidation,
			wantStatusCd: http.StatusBadRequest,
			wantBody:     ErrorResponse{"error": domain.ErrInpValidation.Error()},
		},
		"Unexpected error": {
			rBody:        streetMarketBody{},
			id:           "70ec02cb-0e4a-44cc-b0f7-83c040cb83ea",
			editorErr:    domain.ErrUnexpected,
			wantStatusCd: http.StatusInternalServerError,
			wantBody:     ErrorResponse{"error": domain.ErrUnexpected.Error()},
		},
		"Street Market not founded": {
			rBody:        streetMarketBody{},
			id:           "70ec02cb-0e4a-44cc-b0f7-83c040cb83ea",
			editorErr:    domain.ErrSMNotFound,
			wantStatusCd: http.StatusNotFound,
			wantBody:     ErrorResponse{"error": domain.ErrSMNotFound.Error()},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			editorMock := &stubStreetMarketEditor{
				edit: func(ctx context.Context, id domain.SMID, inp domain.StreetMarketEditInput) error {
					return tc.editorErr
				},
			}

			var body bytes.Buffer
			if err := json.NewEncoder(&body).Encode(tc.rBody); err != nil {
				t.Fatal(err)
			}

			path := fmt.Sprintf("/street_market/%s", tc.id)
			req, err := http.NewRequest(http.MethodPatch, path, &body)
			if err != nil {
				t.Fatal(err)
			}

			h := NewStreetMarketEditHandler(editorMock)
			rr := httptest.NewRecorder()
			r := mux.NewRouter()
			r.HandleFunc("/street_market/{street-market-id}", h.Handle)
			r.ServeHTTP(rr, req)

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
		id := uuid.NewString()
		path := fmt.Sprintf("/street_market/%s", id)
		req, err := http.NewRequest(http.MethodPatch, path, strings.NewReader("body"))
		if err != nil {
			t.Fatal(err)
		}

		h := NewStreetMarketEditHandler(&stubStreetMarketEditor{})
		rr := httptest.NewRecorder()
		r := mux.NewRouter()
		r.HandleFunc("/street_market/{street-market-id}", h.Handle)
		r.ServeHTTP(rr, req)

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
