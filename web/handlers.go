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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	encoder.Encode(data)
}

func newPreacherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if decoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := new(struct {
		Name string
	})
	err := decoder.Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if data.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	preacher, err := db.NewPreacher(data.Name)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		encoder.Encode(map[string]string{"error": err.Error()})
	} else {
		encoder.Encode(preacher)
	}
}
