package main

import (
	"github.com/agocs/quizster/models"
	_ "github.com/agocs/quizster/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.SaveAccount(models.Account{1, "Bob", "Jones", 2014, "bjones@northwestern.edu"})
	beego.Run()
}
