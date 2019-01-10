package controllers

import (
	"CMS/models"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)

type TestPingController struct {
	beego.Controller
}

func (this *TestPingController) PingIPs() {
	recordTime := this.GetString("time")
	beego.Debug(recordTime)

	t := time.Now()
	d, _ := time.ParseDuration("-24h")
	d1 := t.Add(d)
	beego.Debug(d1.Format("20060102"))
	// time.Now().Format("20060102")   date
	// time.Now().Format("1504")   time

	this.ServeJSON()
}

func (this *TestPingController) ReportPingIPs() {

	date1 := this.GetString("date1")
	date2 := this.GetString("date2")
	dateTitle := date1[0:4] + "/" + date1[4:6] + "/" + date1[6:8]

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

	this.Data["Date"] = dateTitle
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

	this.TplName = "test_ping.tpl"
}