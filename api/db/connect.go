package db

import (
	"log"
	"time"

	"github.com/berryhill/gf-api/api/config"

	"gopkg.in/mgo.v2"
)


var Session *mgo.Session  // Session stores mongo session
var Conf *config.Config

func Connect() {

	Conf = config.MakeConfig()

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{Conf.MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: Conf.AuthDatabase,
		Username: Conf.AuthUserName,
		Password: Conf.AuthPassword,
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
