package db

import (
	"log"
	"time"

	"github.com/berryhill/gf-api/api/config"

	"gopkg.in/mgo.v2"
)


var Session *mgo.Session  // Session stores mongo session

func Connect() {

	conf := config.MakeConfig()

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: conf.AuthDatabase,
		Username: conf.AuthUserName,
		Password: conf.AuthPassword,
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
