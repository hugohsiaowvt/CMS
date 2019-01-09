package controllers

import (
	m "CMS/models"
	"fmt"
	"github.com/astaxie/beego"
)

type DashBoardController struct {
	beego.Controller
}

type ShowPageController struct {
	beego.Controller
}

func (c *DashBoardController) Get() {
	Pages:= m.GeneratePages();
	c.Data["Pages"]=Pages
	c.TplName = "dashboard.tpl"

	//內容頁面
	/*c.Layout = "dashboard.tpl"
	page := c.GetString("page")
	if page == "" {
		c.TplName = "testPage1.tpl"
	} else {
		fmt.Printf("\nkey:%s\n", page)
		c.TplName = page
	}*/

}

func (this *ShowPageController) ShowPage() {
	//this.TplName = "dashboard.tpl"
	//內容頁面
	//this.Layout = "dashboard.tpl"
	page := this.GetString("page")
	if page == "" {
		this.TplName = "testPage1.tpl"
	} else {
		fmt.Printf("\nkey:%s\n", page)
		this.TplName = page
	}
}
