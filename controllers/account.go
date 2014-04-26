package controllers

import (
	"github.com/agocs/quizster/models"
	"github.com/astaxie/beego"
	"strconv"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Get() {

	var acct models.Account

	if this.Ctx.Input.Param(":id") != "" {
		acctid, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		acct = models.FindAccount(acctid)
		this.Data["fname"] = acct.Fname
		this.Data["lname"] = acct.Lname
		this.Data["gradyear"] = acct.Gradyear
		this.Data["email"] = acct.Email
		this.TplNames = "show_account.tpl"
	} else {
		this.Data["Website"] = "beego.me"
		this.Data["Email"] = "astaxie@gmail.com"
		this.TplNames = "account.tpl"
	}

}
