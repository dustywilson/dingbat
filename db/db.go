package db

import (
	"os"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Database object
type Database struct {
	*mgo.Database
}

// ObjectID is a bson.ObjectId
type ObjectID bson.ObjectId

var dbAddress = "127.0.0.1"
var dbName = "dingbat"

// ErrNotFound is an mgo.ErrNotFound
var ErrNotFound = mgo.ErrNotFound

// Connect connects to the database
func Connect() *Database {
	if envDBAddress := os.Getenv("MONGO_PORT"); envDBAddress != "" {
		dbAddress = strings.Replace(envDBAddress, "tcp://", "", 1)
	}

	if envDBName := os.Getenv("DB_NAME"); envDBName != "" {
		dbName = envDBName
	}

	var err error
	session, err := mgo.Dial(dbAddress)
	if err != nil {
		panic(err)
	}
	database := session.DB(dbName)
	return &Database{database}
}

// Close closes the database connection
func (d *Database) Close() {
	d.Session.Close()
}

// NewObjectID returns an ObjectID
func NewObjectID() bson.ObjectId {
	return bson.NewObjectId()
}

// NewID returns a new hex-encoded ObjectID
func NewID() string {
	return NewObjectID().Hex()
}
