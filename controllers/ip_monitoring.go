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
	CategoryId    int		`json:"category_id"`
	Id            int64		`json:"id"`
	Title         string	`json:"title"`
	Ip            string	`json:"ip"`
	Monitor_group string	`json:"monitor_group"`
	Type          int		`json:"action"`
}

var Ips []MonitoringItem

func (this *IpMonitoringController) AddIPMonitoring() {

	categoryId, _ := this.GetInt("category_id")
	title := this.GetString("title")
	ip := this.GetString("ip")
	t, _ := this.GetInt("type")

	res := &ResponseStatus{}
	res.Status = -1

	if ip == "" {

	} else {

		o := orm.NewOrm()

		categoryData := &models.TestPingCategory{}

		if err := models.GetCategory(o, categoryData, categoryId); err != nil {

		}

		if id, err := models.AddIPMonitoring(o, categoryId, title, ip, t); err != nil {
			res.Msg = "資料庫錯誤！"
		} else {
			mystruct := MonitoringItem{
				CategoryId:    categoryId,
				Id:            id,
				Title:         title,
				Ip:            ip,
				Monitor_group: categoryData.Category,
				Type:          t,
			}
			Ips = append(Ips,mystruct)

			res.Status = 1
			res.Ext = Ips
		}


	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) EditIPMonitoring() {
	
}

func (this *IpMonitoringController) DelIPMonitoring() {

	categoryId, _ := this.GetInt("category_id")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()

	if err := models.DelIPMonitoring(o, categoryId); err != nil {
		res.Msg = "資料庫錯誤！"
	} else {
		res.Status = 1
		res.Ext = categoryId
	}

	this.Data["json"] = res
	this.ServeJSON()

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
		mystruct := MonitoringItem{
			CategoryId: v.CategoryId,
			Title: v.Item,
			Ip: v.Ip,
			Monitor_group: v.Category,
			Type: v.Type,
		}
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

	res := &ResponseStatus{}

	var categoryId int
	var category, item, ip string

	res.Status = -1

	if isToday := CheckIsToday(date); isToday {

		o := orm.NewOrm()

		allData := []models.TestPingData{}
		if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
			beego.Debug(err)
			res.Msg = "資料庫錯誤！"
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
			res.Msg = "資料庫錯誤！"
		} else {
			res.Status = 1
		}
	} else {
		res.Msg = "無法新增非當天報表資料！"
	}
	fmt.Print(res)
	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) EditPingResult() {

	id, _ := strconv.Atoi(this.GetString("result_id"))
	status, _ := strconv.Atoi(this.GetString("status"))

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()

	result := &models.TestPingData{}

	if err := models.GetPingResult(o, result, id); err != nil {
		beego.Debug(err)
		res.Msg = "資料庫錯誤！"
	} else {
		date := result.Date
		if isToday := CheckIsToday(date); isToday {
			if err := models.EditPingResult(o, id, status); err != nil {
				beego.Debug(err)
				res.Msg = "資料庫錯誤！"
			} else {
				res.Status = 1
			}
		} else {
			// 非當日
			res.Msg = "無法修改非當天報表資料！"
		}
	}
	this.Data["json"] = res
	this.ServeJSON()

}