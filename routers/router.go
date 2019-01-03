package routers

import (
	"CMS/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

func init() {
	//var FilterUser = func(ctx *context.Context) {
	//	_, ok := ctx.Input.Session("uid").(int)
	//	if !ok && ctx.Request.RequestURI != "/" {
	//		ctx.Redirect(302, "/")
	//	}
	//}
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
    beego.Router("/", &controllers.LoginController{})
	beego.Router("/dashboard", &controllers.DashBoardController{})
	beego.Router("/test", &controllers.TestController{})
}
