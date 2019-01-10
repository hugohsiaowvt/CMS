package controllers

import (
	"CMS/models"
	"CMS/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

type TestPingController struct {
	beego.Controller
}

var waitGroup = new(sync.WaitGroup)

func (this *TestPingController) PingIPs() {
	recordTime := this.GetString("time")

	t := time.Now().Local()
	date := t.Format("2006-01-02")

	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if _, err := models.GetAllTestPingData(o, &allData); err != nil {
		this.ServeJSON()
	}

	for _, v := range allData {
		waitGroup.Add(1)
		go ping(v, date, recordTime)
	}

	waitGroup.Wait()

	this.ServeJSON()
}

func ping(data models.TestPingData, date, recordTime string) {
	o := orm.NewOrm()
	status := -1
	for i := 0; i < 5; i++ {
		if alive := utils.Ping(data.Ip); alive {
			status = 1
		}
	}

	if err := models.InsertTestPingResult(o, data, date, recordTime, status); err != nil {

	}

	waitGroup.Done()
}

func (this *TestPingController) ReportPingIPs() {

	date1 := this.GetString("date1")
	date2 := this.GetString("date2")

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

	this.Data["Date"] = date1
	this.Data["Data"] = allData
	this.Data["Count"] = count
	this.Data["Result"] = resultData
	this.Data["Times"] = [] string{
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

	this.TplName = "report_ping_ips.tpl"
}