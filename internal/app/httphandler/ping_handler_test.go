package httphandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPingHandler_Handle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	ph := NewPingHandler()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ph.Handle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `"pong"`
	got := rr.Body.String()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
	}
}
