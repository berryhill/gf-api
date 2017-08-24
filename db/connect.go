package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)


var Session *mgo.Session  // Session stores mongo session

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = ""
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "test"
)

func Connect() {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	Session = session
}

func Clone() *mgo.Session {
	
	session := Session.Clone()
	defer session.Close()

	return session
}
