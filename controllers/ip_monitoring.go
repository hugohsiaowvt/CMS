package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IpMonitoringController struct {
	beego.Controller
}

type MonitoringItem struct {
	Title 			string `json:"title"`
	Ip	  			string `json:"ip"`
	Monitor_group	string `json:"monitor_group"`
}

func (this *IpMonitoringController) AddIPMonitoring() {
	fmt.Printf("test3\n")
	//TODO 處理寫入DB 後將成功訊息回傳
	mystruct := MonitoringItem{ Title:"1234",Ip:"12132",Monitor_group:"2rgss"}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}

func (this *IpMonitoringController) DelIPMonitoring() {

}