package models

import (
	"github.com/agocs/quizster/database"
	_ "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Topic struct {
	Id   int
	Name string
}

func FindTopic(id int) Topic {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("topics")

	t := Topic{}
	err := c.Find(bson.M{"id": id}).One(&t)
	if err != nil {
		log.Printf("unable to find topic %d", id)
		log.Panic(err.Error())
	}
	log.Printf("Read Topic %d from database", id)

	return t
}

func FindAllTopics() []Topic {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("topics")

	t := []Topic{}
	err := c.Find(nil).All(&t)
	if err != nil {
		log.Printf("unable to find all topics -- %s", err.Error())
	}

	return t
}

func SaveTopic(t Topic) bool {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("topics")
	err := c.Insert(t)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Printf("Inserted topic %d", t.Id)

	return true
}
