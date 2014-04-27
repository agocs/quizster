package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Account struct {
	Id       int
	Fname    string
	Lname    string
	Gradyear int
	Email    string
}

func FindAccount(id int) Account {
	session, err := mgo.Dial("107.170.101.48")
	if err != nil {
		log.Println("Unable to connect to database.")
		log.Panic(err.Error())
	}

	defer session.Close()

	c := session.DB("quizster").C("accounts")

	a := Account{}
	err = c.Find(bson.M{"id": id}).One(&a)
	if err != nil {
		log.Printf("unable to find user %d", id)
		log.Panic(err.Error())
	}
	log.Printf("Read Account %d from database", id)

	return a
}

func SaveAccount(a Account) bool {
	session, err := mgo.Dial("107.170.101.48")
	if err != nil {
		log.Println("Unable to connect to database.")
		log.Panic(err.Error())
		panic(err)
	}
	defer session.Close()

	c := session.DB("quizster").C("accounts")
	err = c.Insert(a)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Printf("Inserted account %d", a.Id)

	return true
}
