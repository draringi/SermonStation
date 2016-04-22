package web

import (
	"github.com/draringi/SermonStation/audio"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var audioManager *audio.Manager

func StartServer(AudioManager *audio.Manager) {
	router := getRouter()
	audioManager = AudioManager
	http.Handle("/api/", router)
	go func() {
		for {
			log.Println(http.ListenAndServe(":8080", nil))
		}
	}()
}
