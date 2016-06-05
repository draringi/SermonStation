package web

import (
	"encoding/json"
	"github.com/draringi/SermonStation/db"
	"log"
	"net/http"
	"path"
	"time"
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

func liveRecordingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if decoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := new(struct {
		Command  string
		Preacher int
		Keyword  string
		Title    string
	})
	err := decoder.Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	switch data.Command {
	default:
		w.WriteHeader(http.StatusBadRequest)
	case "new":
		preacher, err := db.PreacherByID(data.Preacher)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			encoder.Encode(map[string]string{"error": err.Error()})
			return
		}
		startTime := time.Now()
		webPath := pathGenerator(preacher.Name(), data.Keyword, startTime)
		abdPath := path.Clean(baseDir + webPath)
		dbRec, err := db.NewRecording(startTime, webPath, preacher)
		if err != nil {
			w.WriteHeader(http.http.StatusInternalServerError)
			log.Println(err)
			return
		}
		rec, err := audioManager.NewRecording(absPath)
		if err != nil {
			w.WriteHeader(http.StatusPaymentRequired)
			encoder.Encode(map[string]string{"error": err.Error()})
			return
		}
		encoder.Encode(rec.Status())
		return
	case "start":
		rec := audioManager.Recording()
		if rec == nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(map[string]string{"error": "No Active Recording"})
			return
		}
		err := rec.Start()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		encoder.Encode(rec.Status())
		return
	case "stop":
		rec := audioManager.Recording()
		if rec == nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(map[string]string{"error": "No Active Recording"})
			return
		}
		rec.Stop()
		encoder.Encode(rec.Status())
		return
	case "get":
		rec := audioManager.Recording()
		if rec == nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(map[string]string{"error": "No Active Recording"})
			return
		}
		encoder.Encode(rec)
		return
	}
}
