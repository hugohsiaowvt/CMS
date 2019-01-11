package controllers

import (
	"CMS/conf"
	"fmt"
	"CMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
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

func (this *IpMonitoringController) GetIPList() {
	o := orm.NewOrm()
	allData := []models.TestPingData{}
	if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
		return
	}
	Ips = Ips[:0]
	fmt.Printf("count:%d\n", len(Ips))
	for _, v := range allData {
		mystruct := MonitoringItem{ Title:v.Item,Ip:v.Ip,Monitor_group:v.Category}
		Ips = append(Ips,mystruct)
	}

	this.Data["json"] = &Ips
	this.ServeJSON()
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

	isToday := true
	if !CheckIsToday(date1) {
		isToday = false
	}


	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if isToday {
		if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
			beego.Debug(err)
		}
	} else {
		if _, err := models.GetPreviousPingData(o, &allData, date1, date2); err != nil {
			beego.Debug(err)
		}
	}

	countData := []models.TestPingCategoryCount{}
	if isToday {
		if _, err := models.GetBaseCategoryCount(o, &countData); err != nil {
			beego.Debug(err)
		}
	} else {
		if _, err := models.GetPreviousCategoryCount(o, &countData, date1, date2); err != nil {
			beego.Debug(err)
		}
	}

	resultData := []models.TestPingResultData{}
	if _, err := models.GetTestPingResultByDate(o, &resultData, date1, date2); err != nil {
		beego.Debug(err)
	}

	count := make(map[int]int)
	for _, v := range countData {
		count[v.Id] = v.Count
	}

	fmt.Printf("resultData:%d\n", len(resultData))
	times := conf.PING_TIME

	this.Data["json"] = &PingData{ Date: date1 , AllData:allData, Count:count, Result:resultData , Times:times ,TestPlanCase:conf.TEST_PLAN_TIME  }
	this.ServeJSON()
}

func (this *IpMonitoringController) AddPingResult() {

	itemId, _ := strconv.Atoi(this.GetString("item_id"))
	date := this.GetString("date")
	time := this.GetString("time")
	status, _ := strconv.Atoi(this.GetString("status"))

	var categoryId int
	var category, item, ip string

	beego.Debug(itemId, date, time, status)

	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
		beego.Debug(err)
	}

	for _, v := range allData {
		if v.ItemId == itemId {
			categoryId = v.CategoryId
			category = v.Category
			item = v.Item
			ip = v.Ip
		}
	}

	if err := models.AddPingResult(o, categoryId, itemId, status, date, time, category, item, ip); err != nil {
		beego.Debug(err)
	}

	this.ServeJSON()

}

func (this *IpMonitoringController) EditPingResult() {

}