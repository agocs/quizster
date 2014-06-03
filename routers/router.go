package routers

import (
	"github.com/agocs/quizster/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/account", &controllers.AccountController{})
	beego.Router("/account/:id", &controllers.AccountController{})
	beego.Router("/question/:id", &controllers.QuestionController{})
	beego.Router("/login", &controllers.LoginController{})
}
