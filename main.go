package main

import (
	"github.com/agocs/quizster/models"
	_ "github.com/agocs/quizster/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.SaveAccount(models.Account{1, "Bob", "Jones", 2014, "bjones@northwestern.edu"})
	models.SaveTopic(models.Topic{1, "Potpouri"})
	models.SaveQuestion(models.Question{1, 1, "What what, in the butt?"})
	models.SaveAnswerChoice(models.AnswerChoice{1, 1, "Right answer", true})
	models.SaveAnswerChoice(models.AnswerChoice{2, 1, "Wrong 1", false})
	models.SaveAnswerChoice(models.AnswerChoice{3, 1, "Wrong 2", false})
	models.SaveAnswerChoice(models.AnswerChoice{4, 1, "You dummy", false})
	beego.Run()
}
