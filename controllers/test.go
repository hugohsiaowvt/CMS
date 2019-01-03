package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

//func (this *TestController) Prepare() {
//	v := this.GetSession("account")
//	if v == nil {
//		this.Redirect("/", 302)
//	}
//}

func (this *TestController) Get() {
	this.TplName = "test.tpl"
}
