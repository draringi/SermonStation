package main

import (
	"github.com/draringi/SermonStation/audio"
	"github.com/draringi/SermonStation/db"
	"github.com/draringi/SermonStation/web"
	"github.com/gordonklaus/portaudio"
	"log"
	"os"
	"os/signal"
	"os/user"
)

func main() {
	defer portaudio.Terminate()
	audioManager, err := audio.NewManager()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	dbUser := os.Getenv("SERMON_USER")
	if dbUser == "" {
		currentUser, err := user.Current()
		if err != nil {
			log.Printf("ERROR: %v, Reverting to default username\n", err)
			dbUser = "sermons"
		} else {
			dbUser = currentUser.Username
		}

	}
	dbDatabase := os.Getenv("SERMON_DB")
	if dbDatabase == "" {
		dbDatabase = "sermons"
	}
	log.Printf("Connecting to DB %s as User %s\n", dbDatabase, dbUser)
	err = db.ConnectToDatabase(dbUser, dbDatabase)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return
	}
	log.Println("Initiating Web Server")
	web.StartServer(audioManager)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	for {
		select {
		case <-sig:
			log.Println("Signal Recieved. Shutting Down.")
			return
		default:
		}
	}
}
