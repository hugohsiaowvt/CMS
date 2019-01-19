package controllers

import (
	"CMS/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type ReportMonitoringController struct {
	beego.Controller
}

type ReportData struct {
	Date			string
	AllData 		[]models.MonitoringSchema
	Count 			map[int]int
	Result 			[]models.MonitoringResultData
}

func (this *ReportMonitoringController) GenerateMonitoringRecord() {

	t := time.Now().Local()
	date := t.Format("2006-01-02")

	data := []models.MonitoringSchema{}

	o := orm.NewOrm()

	if _, err := models.GetBaseAllReportMonitoringSchema(o, &data); err != nil {

	}

	reportResult := []models.ReportResult {}

	// 加入日期
	for _, v := range data {
		reportResult = append(reportResult, models.ReportResult{
			CategoryId: v.CategoryId,
			ItemId: v.ItemId,
			Category: v.Category,
			Item: v.Item,
			Date: date,
		})
	}

	if err := models.InsertMonitoringResult(o, date, reportResult); err != nil {

	}

	this.ServeJSON()

}

func (this *ReportMonitoringController) GetData() {

	date := this.GetString("date")

	if date == "" {
		t := time.Now().Local()
		date = t.Format("2006-01-02")
	}

	o := orm.NewOrm()

	allData := []models.MonitoringSchema{}
	if _, err := models.GetBaseAllReportMonitoringSchema(o, &allData); err != nil {
		beego.Debug(err)
	}

	countData := []models.MonitoringCategoryCount{}
	if _, err := models.GetBaseReportMonitoringCategoryCount(o, &countData); err != nil {
		beego.Debug(err)
	}

	resultData := []models.MonitoringResultData{}
	if _, err := models.GetReportMonitoringResultByDate(o, &resultData, date); err != nil {
		beego.Debug(err)
	}

	count := make(map[int]int)
	for _, v := range countData {
		count[v.Id] = v.Count
	}

	fmt.Printf("resultData:%d\n", len(resultData))

	this.Data["json"] = &ReportData{ Date: date, AllData: allData, Count: count, Result:resultData }
	this.ServeJSON()

}