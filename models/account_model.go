package models

import (
	"github.com/agocs/quizster/database"
	_ "labix.org/v2/mgo"
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
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("accounts")

	a := Account{}
	err := c.Find(bson.M{"id": id}).One(&a)
	if err != nil {
		log.Printf("unable to find user %d", id)
		log.Panic(err.Error())
	}
	log.Printf("Read Account %d from database", id)

	return a
}

func SaveAccount(a Account) bool {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("accounts")
	err := c.Insert(a)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Printf("Inserted account %d", a.Id)

	return true
}

func GetCurrentAccount() Account {
	//#TODO: Write later
	return FindAccount(1)
}
