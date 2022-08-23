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
	"github.com/gorilla/mux"
)

type stubStreetMarketEraser struct {
	deleteInp domain.SMID
	delete    func(context.Context, domain.SMID) error
}

func (s *stubStreetMarketEraser) Delete(ctx context.Context, id domain.SMID) error {
	s.deleteInp = id
	return s.delete(ctx, id)
}

func TestStreetMarketDeleteHandler_Handle(t *testing.T) {
	id := domain.SMID("cdd2028a-fd0b-4734-97e7-ef2e57e9009b")

	eraserMock := &stubStreetMarketEraser{
		delete: func(ctx context.Context, s domain.SMID) error {
			return nil
		},
	}

	path := fmt.Sprintf("/street_market/%s", id)
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	h := NewStreetMarketDeleteHandler(eraserMock)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/street_market/{street-market-id}", h.Handle)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expect status code %v, got %v", status, http.StatusNoContent)
	}

	if id != eraserMock.deleteInp {
		t.Errorf(
			"street market eraser delete receive a unexpected street market id, want %s, got %s",
			id,
			eraserMock.deleteInp,
		)
	}
}

func TestStreetMarketDeleteHandler_Handle_Error(t *testing.T) {
	testCases := map[string]struct {
		id           string
		eraserErr    error
		wantStatusCd int
		wantBody     ErrorResponse
	}{
		"Unexpected error": {
			id:           "5e5e1905-282f-4038-b8c9-bc1719174892",
			eraserErr:    domain.ErrUnexpected,
			wantStatusCd: http.StatusInternalServerError,
			wantBody:     ErrorResponse{"error": domain.ErrUnexpected.Error()},
		},
		"Invalid person id": {
			id:           "id",
			eraserErr:    domain.ErrInpValidation,
			wantStatusCd: http.StatusBadRequest,
			wantBody:     ErrorResponse{"error": domain.ErrInpValidation.Error()},
		},
		"Street Market not founded": {
			id:           "id",
			eraserErr:    domain.ErrSMNotFound,
			wantStatusCd: http.StatusNotFound,
			wantBody:     ErrorResponse{"error": domain.ErrSMNotFound.Error()},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			eraserMock := &stubStreetMarketEraser{
				delete: func(ctx context.Context, s domain.SMID) error {
					return tc.eraserErr
				},
			}

			path := fmt.Sprintf("/street_market/%s", tc.id)
			req, err := http.NewRequest(http.MethodDelete, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			h := NewStreetMarketDeleteHandler(eraserMock)
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
}
