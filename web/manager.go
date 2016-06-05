package web

import (
	"github.com/draringi/SermonStation/audio"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

const (
	defaultBaseDir = "/usr/local/www/sermons/"
)

var baseDir string = os.Getenv("SERMON_BASEDIR")

var audioManager *audio.Manager

func StartServer(AudioManager *audio.Manager) {
	if baseDir == "" {
		baseDir = defaultBaseDir
	}
	router := getRouter()
	audioManager = AudioManager
	http.Handle("/api/", router)
	go func() {
		for {
			log.Println(http.ListenAndServe(":8080", nil))
		}
	}()
}
