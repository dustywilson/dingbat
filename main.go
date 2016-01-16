package main

import (
	"net/http"

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

	http.Handle("/gotalk/", ws)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	panic(http.ListenAndServe(":7654", nil))
}
