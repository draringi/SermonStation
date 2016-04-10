package web

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func StartServer() {
	router := getRouter()
	http.Handle("/", router)
	go func() {
		for {
			log.Println(http.ListenAndServe(":8080", nil))
		}
	}()
}
