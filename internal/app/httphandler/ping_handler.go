package httphandler

import (
	"net/http"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body := "pong"

	respondJSON(w, http.StatusOK, body)
}
