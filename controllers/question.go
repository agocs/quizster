package controllers

import (
	"github.com/agocs/quizster/models"
	"github.com/astaxie/beego"
	"strconv"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) Get() {
	if this.Ctx.Input.Param(":id") != "" {
		qid, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		q := models.FindQuestion(qid)
		t := models.FindTopic(q.TopicId)

		this.Data["topic"] = t.Name
		this.Data["question"] = q.QuestionText
		this.Data["answer1"] = q.RightAnswer
		this.Data["answer2"] = q.WrongAnswer1
		this.Data["answer3"] = q.WrongAnswer2
		this.Data["answer4"] = q.WrongAnswer3
		this.TplNames = "show_question.tpl"
	} else {
		//show me all the questions
	}
}
