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
	beego.Router("/dashboard/showPage", &controllers.ShowPageController{},"*:ShowPage")
	beego.Router("/test_ping", &controllers.DashBoardController{},"*:TestPing")

	beego.Router("/monitoring/add", &controllers.IpMonitoringController{},"*:AddIPMonitoring")
	beego.Router("/monitoring/edit", &controllers.IpMonitoringController{},"*:EditIPMonitoring")
	beego.Router("/monitoring/del", &controllers.IpMonitoringController{},"*:DelIPMonitoring")
	beego.Router("/monitoring/ping", &controllers.IpMonitoringController{},"*:TestPing")
	beego.Router("/monitoring/ips", &controllers.IpMonitoringController{},"*:GetIPList")
	beego.Router("/pingresult/add", &controllers.IpMonitoringController{},"*:AddPingResult")
	beego.Router("/pingresult/edit", &controllers.IpMonitoringController{},"*:EditPingResult")

	beego.Router("/report/pingips", &controllers.TestPingController{}, "*:ReportPingIPs")
	beego.Router("/api/pingips", &controllers.TestPingController{}, "*:PingIPs")

	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)

}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uname").(string)
	if !ok && ctx.Request.RequestURI != "/" {
		// api和報表不用登入就可以呼叫
		if string(ctx.Request.RequestURI[0:4]) != "/api" && string(ctx.Request.RequestURI[0:7]) != "/report" {
			ctx.Redirect(302, "/")
		}
	}
}