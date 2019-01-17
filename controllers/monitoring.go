package controllers

import (
	"CMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type MonitoringController struct {
	beego.Controller
}

func (this *MonitoringController) GenerateMonitoringRecord() {

	t := time.Now().Local()
	date := t.Format("2006-01-02")

	data := []models.MonitoringSchema{}

	o := orm.NewOrm()

	if _, err := models.GetBaseAllMonitoringSchema(o, &data); err != nil {

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