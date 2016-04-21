package web

import (
	"encoding/json"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder.Encode(audioManager.Status())
}

func preachers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
}
