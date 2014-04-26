package controllers

import (
	//"github.com/agocs/quizster/models"
	"github.com/astaxie/beego"
	//"strconv"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	if this.Ctx.Input.Param(":id") != "" {
		//grab all the questions with this topicID
		this.TplNames = "show_topic.tpl"
	} else {
		//grab all the topics
	}
}
