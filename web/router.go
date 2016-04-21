package web

import (
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", statusHandler)
	return r
}
