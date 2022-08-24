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
	delete    func(context.Context, domain.SMID) *domain.Error
}

func (s *stubStreetMarketEraser) Delete(ctx context.Context, id domain.SMID) *domain.Error {
	s.deleteInp = id
	return s.delete(ctx, id)
}

func TestStreetMarketDeleteHandler_Handle(t *testing.T) {
	id := domain.SMID("cdd2028a-fd0b-4734-97e7-ef2e57e9009b")

	eraserMock := &stubStreetMarketEraser{
		delete: func(ctx context.Context, s domain.SMID) *domain.Error {
			return nil
		},
	}

	path := fmt.Sprintf("/street_market/%s", id)
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	h := NewStreetMarketDeleteHandler(eraserMock, &stubLogger{})
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
		eraserErr    *domain.Error
		wantStatusCd int
		wantBody     ErrorResponse
	}{
		"Unexpected error": {
			id:           "5e5e1905-282f-4038-b8c9-bc1719174892",
			eraserErr:    &domain.Error{Kind: domain.UnexpectedErrKd, Msg: "Unexpected"},
			wantStatusCd: http.StatusInternalServerError,
			wantBody:     ErrorResponse{"error": "Unexpected"},
		},
		"Invalid person id": {
			id:           "id",
			eraserErr:    &domain.Error{Kind: domain.InpValidationErrKd, Msg: "Error"},
			wantStatusCd: http.StatusBadRequest,
			wantBody:     ErrorResponse{"error": "Error"},
		},
		"Street Market not founded": {
			id:           "id",
			eraserErr:    &domain.Error{Kind: domain.SMNotFoundErrKd, Msg: "SM not found"},
			wantStatusCd: http.StatusNotFound,
			wantBody:     ErrorResponse{"error": "SM not found"},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			eraserMock := &stubStreetMarketEraser{
				delete: func(ctx context.Context, s domain.SMID) *domain.Error {
					return tc.eraserErr
				},
			}

			path := fmt.Sprintf("/street_market/%s", tc.id)
			req, err := http.NewRequest(http.MethodDelete, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			h := NewStreetMarketDeleteHandler(eraserMock, &stubLogger{})
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
