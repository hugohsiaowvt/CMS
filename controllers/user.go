package controllers

import (
	"CMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:aS!23456@tcp(173.248.224.95:3306)/CMS?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
}

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	if this.Ctx.Input.IsGet() {
		// 獲取 session
		userName := this.GetSession("uname")
		userPwd := this.GetSession("upwd")

		//log.Println("user:" + userName.(string))

		_, nameOk := userName.(string)
		_, pwdOk := userPwd.(string)
		if nameOk && pwdOk {
			// 重定向
			this.Redirect("dashboard", 302)
		} else {
			// 获取 cookie
			this.Data["uname"] = this.Ctx.GetCookie("uname")
			this.Data["upwd"] = this.Ctx.GetCookie("upwd")
			this.TplName = "login.tpl"
		}
	} else {
		userName := this.GetString("uname")
		userPwd := this.GetString("upwd")

		o := orm.NewOrm()
		user := &models.Users{}
		query := &models.Users{
			Account: userName,
			Password: userPwd,
		}

		// insert
		models.GetUser(o, user, query)

		if user.Status == 1 {
			// 设置 cookie
			this.Ctx.SetCookie("uname", userName)
			this.Ctx.SetCookie("upwd", userPwd)
			// 设置 session
			this.SetSession("uname", userName)
			this.SetSession("upwd", userPwd)
			this.Redirect("dashboard", 302)
		} else {
			this.Data["uname"] = this.Ctx.GetCookie("uname")
			this.Data["upwd"] = this.Ctx.GetCookie("upwd")
			this.Data["isError"] = true
			this.TplName = "login.tpl"
		}

	}
}

func (this *UserController) Exit() {
	this.DelSession("uname")
	this.DelSession("upwd")
	this.Data["json"] = nil
	this.Redirect("", 302)
}