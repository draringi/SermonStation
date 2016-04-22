package web

import (
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.HandleFunc("/status/", statusHandler)
	r.HandleFunc("/preachers/", preachersListHandler)
	return r
}
