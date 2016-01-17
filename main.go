package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dustywilson/dingbat/client"
	"github.com/dustywilson/dingbat/db"
	_ "github.com/dustywilson/dingbat/person"
	"github.com/rsms/gotalk"
)

var database *db.Database

func main() {
	database = db.Connect()
	defer database.Close()

	ws := gotalk.WebSocketHandler()
	ws.OnAccept = onAccept

	wwwPath := "./www"
	if envWWWPath := os.Getenv("WWW_PATH"); envWWWPath != "" {
		wwwPath = envWWWPath
	}

	listenAddr := ":7654"
	if envListenAddr := os.Getenv("WWW_PATH"); envListenAddr != "" {
		listenAddr = envListenAddr
	}

	http.Handle("/gotalk/", ws)
	http.Handle("/", http.FileServer(http.Dir(wwwPath)))
	panic(http.ListenAndServe(listenAddr, nil))
}

func onAccept(s *gotalk.Sock) {
	c := client.NewClient()
	s.UserData = c
	fmt.Printf("New accepted connection.  [%s]\n", c.ID)
}
