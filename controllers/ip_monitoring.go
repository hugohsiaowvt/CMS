package controllers

import (
	"fmt"
	"CMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type IpMonitoringController struct {
	beego.Controller
}

type MonitoringItem struct {
	Title 			string `json:"title"`
	Ip	  			string `json:"ip"`
	Monitor_group	string `json:"monitor_group"`
}

var Ips []MonitoringItem

func (this *IpMonitoringController) AddIPMonitoring() {
	fmt.Printf("test3\n")
	//TODO 處理寫入DB 後將成功訊息回傳
	group := this.GetString("group")
	title := this.GetString("title")
	ip := this.GetString("ip")
	if group=="" || title == "" || ip == "" {

	} else {
		mystruct := MonitoringItem{ Title:title,Ip:ip,Monitor_group:group}
		Ips = append(Ips,mystruct)
	}

	this.Data["json"] = &Ips
	this.ServeJSON()

}

func (this *IpMonitoringController) DelIPMonitoring() {

}

type PingData struct {
	Date			string
	AllData 		[]models.TestPingData
	Count 			map[int]int
	Result 			[]models.TestPingResultData
	Times			[] string
	TestPlanCase	[]string
}

func (this *IpMonitoringController) TestPing() {

	date1 := this.GetString("date")

	if date1 == "" {
		t := time.Now().Local()
		date1 = t.Format("2006-01-02")
	}

	now, _ := time.Parse("2006-01-02", date1)
	nextDay, _ := time.ParseDuration("24h")
	now = now.Add(nextDay)
	date2 := now.Format("2006-01-02")
	beego.Debug(date1, date2)

	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if _, err := models.GetAllTestPingData(o, &allData); err != nil {
		return
	}

	countData := []models.TestPingCategoryCount{}
	if _, err := models.GetCategoryCount(o, &countData); err != nil {
		return
	}

	resultData := []models.TestPingResultData{}
	if _, err := models.GetTestPingResultByDate(o, &resultData, date1, date2); err != nil {
		return
	}

	count := make(map[int]int)
	for _, v := range countData {
		count[v.Id] = v.Count
	}

	fmt.Printf("resultData:%d\n", len(resultData))
	times := [] string{
		"1830",
		"1945",
		"2015",
		"2130",
		"2245",
		"2315",
		"0030",
		"0145",
		"0215",
		"0330",
		"0445",
		"0555",
	}

	this.Data["json"] = &PingData{ Date: date1 , AllData:allData, Count:count, Result:resultData , Times:times ,TestPlanCase:models.GetTestPlanTime()  }
	this.ServeJSON()
}