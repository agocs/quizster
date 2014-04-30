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
		ac := models.FindAnswersForQuestion(qid)

		this.Data["topic"] = t.Name
		this.Data["question"] = q.QuestionText

		answerchoices := ""
		for _, answerchoice := range ac {
			answerchoices = answerchoices + answerchoice.Text + "\n"
		}
		this.Data["answerChoices"] = answerchoices
		this.TplNames = "show_question.tpl"

	} else {
		//show me all the questions
	}
}
