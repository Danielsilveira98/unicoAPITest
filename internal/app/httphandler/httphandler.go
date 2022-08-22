package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse map[string]interface{}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println(err) // TODO log here
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(response)); err != nil {
		fmt.Println(err) // TODO log here
	}
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, ErrorResponse{"error": message})
}
