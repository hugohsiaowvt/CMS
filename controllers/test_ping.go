package controllers

import (
	"CMS/conf"
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
	if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
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

	isToday := true
	if !CheckIsToday(date1) {
		isToday = false
	}

	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if isToday {
		if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
			return
		}
	} else {
		if _, err := models.GetPreviousPingData(o, &allData, date1, date2); err != nil {
			return
		}
	}

	countData := []models.TestPingCategoryCount{}
	if isToday {
		if _, err := models.GetBaseCategoryCount(o, &countData); err != nil {
			return
		}
	} else {
		if _, err := models.GetPreviousCategoryCount(o, &countData, date1, date2); err != nil {
			return
		}
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
	this.Data["Times"] = conf.PING_TIME

	this.TplName = "report_ping_ips.tpl"
}