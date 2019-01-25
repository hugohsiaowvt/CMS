package controllers

import (
	"CMS/conf"
	"CMS/models"
	"encoding/json"
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
			Note: "",
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
	if _, err := models.GetPreviousReportMonitoringSchema(o, &allData, date); err != nil {
		beego.Debug(err)
	}

	countData := []models.MonitoringCategoryCount{}
	if _, err := models.GetPreviousReportMonitoringCategoryCount(o, &countData, date); err != nil {
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

func (this *ReportMonitoringController) EditReportStatus() {

	id, _ := this.GetInt("result_id")
	status, _ := this.GetInt("status")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	data := models.EditStatusAfter{
		Id: id,
		Status: status,
	}

	b, _ := json.Marshal(data)

	userName := this.GetSession("uname")
	if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["EDIT_MONITORING_RESULT"], userName,"", string(b)); err != nil {
		res.Msg = err.Error()
		o.Rollback()
	} else {
		if err := models.EditReportMonitoringStatus(o, id, status); err != nil {
			res.Msg = err.Error()
		} else {
			res.Status = 1
			o.Commit()
		}
	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *ReportMonitoringController) EditReportNote() {

	id, _ := this.GetInt("result_id")
	beego.Debug(id)
	note := this.GetString("note")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	data := &models.MonitoringResultData{}
	if err := models.GetReportNonitoringResultById(o, data, id); err != nil {

	}

	bdata := models.EditNoteBefore {
		Id: id,
		Note: data.Note,
	}
	b, _ := json.Marshal(bdata)


	adata := models.EditNoteAfter {
		Id: id,
		Note: note,
	}
	a, _ := json.Marshal(adata)

	userName := this.GetSession("uname")
	if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["EDIT_MONITORING_RESULT"], userName, string(b), string(a)); err != nil {
		res.Msg = err.Error()
		o.Rollback()
	} else {
		if err := models.EditReportMonitoringNote(o, id, note); err != nil {
			res.Msg = err.Error()
		} else {
			res.Status = 1
			o.Commit()
		}
	}

	this.Data["json"] = res
	this.ServeJSON()

}