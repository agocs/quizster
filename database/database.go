package database

import (
	"labix.org/v2/mgo"
	"log"
)

func MakeSession() *mgo.Session {
	session, err := mgo.Dial("107.170.101.48")
	if err != nil {
		log.Println("Unable to connect to database.")
		log.Panic(err.Error())
	}
	return session
}
