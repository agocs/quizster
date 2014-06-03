package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type Credentials struct {
	Username string
	Password string
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"

}

func (this *LoginController) Post() {
	request := this.Ctx.Request
	p := make([]byte, request.ContentLength)
	_, err := this.Ctx.Request.Body.Read(p)
	if err != nil {
		panic(err.Error())
	}
	var theseCredentials Credentials
	err1 := json.Unmarshal(p, &theseCredentials)
	if err1 != nil {
		panic(err.Error())
	}
	this.Data["Username"] = theseCredentials.Username
	this.Data["Password"] = theseCredentials.Password
}
