package person

import "github.com/dustywilson/dingbat/db"

var database = db.Connect()
var coll = database.C("person")
