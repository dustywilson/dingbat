package person

import "gopkg.in/mgo.v2/bson"

// Person is a human person
type Person struct {
	ID        bson.ObjectId `bson:"_id"`
	FirstName string
	LastName  string
}
