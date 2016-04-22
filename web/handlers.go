package web

import (
	"encoding/json"
	"github.com/draringi/SermonStation/db"
	"log"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder.Encode(audioManager.Status())
}

func preachersListHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := db.ListPreachers()
	log.Println(data, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder.Encode(data)
}
