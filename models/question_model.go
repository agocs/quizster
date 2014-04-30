package models

import (
	"github.com/agocs/quizster/database"
	_ "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Question struct {
	Id           int
	TopicId      int
	QuestionText string
}

type AnswerChoice struct {
	Id         int
	QuestionId int
	Text       string
	IsCorrect  bool
}

func FindQuestion(id int) Question {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("questions")

	q := Question{}
	err := c.Find(bson.M{"id": id}).One(&q)
	if err != nil {
		log.Printf("unable to find question %d", id)
		log.Panic(err.Error())
	}

	return q
}

func FindAnswersForQuestion(qid int) []AnswerChoice {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("answerchoices")

	a := []AnswerChoice{}
	err := c.Find(bson.M{"questionid": qid}).All(&a)
	if err != nil {
		log.Printf("unable to find answer choices for question %d", qid)
		log.Panic(err.Error())
	}
	log.Printf("Read %d answer choices for question %d from database", len(a), qid)

	return a
}

func SaveQuestion(q Question) bool {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("questions")
	err := c.Insert(q)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Printf("Inserted question %d", q.Id)

	return true
}

func SaveAnswerChoice(a AnswerChoice) bool {
	session := database.MakeSession()
	defer session.Close()

	c := session.DB("quizster").C("answerchoices")
	err := c.Insert(a)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Printf("Inserted answer choice %d", a.Id)

	return true
}
