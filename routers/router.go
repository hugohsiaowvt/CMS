package routers

import (
	"CMS/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.Router("/", &controllers.UserController{}, "*:Login")
	beego.Router("/logout", &controllers.UserController{}, "*:Exit")
	beego.Router("/dashboard", &controllers.DashBoardController{})

	beego.Router("/test_ping", &controllers.DashBoardController{}, "*:TestPing")
	beego.Router("/showPage", &controllers.ShowPageController{},"*:ShowPage")


	beego.Router("/monitoring/add", &controllers.IpMonitoringController{},"*:AddIPMonitoring")
	beego.Router("/monitoring/del", &controllers.IpMonitoringController{},"*:DelIPMonitoring")
	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)

}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uname").(string)
	if !ok && ctx.Request.RequestURI != "/" {
		ctx.Redirect(302, "")
	}
}