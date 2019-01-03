package controllers

import "github.com/astaxie/beego"

type DashBoardController struct {
	beego.Controller
}

func (c *DashBoardController) Get() {
	c.TplName = "dashboard.tpl"
}