package controllers

import (
	"CMS/conf"
	"encoding/json"
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

func (this *IpMonitoringController) GetCategoryName() {

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	category := &[]models.TestPingCategory{}

	if _, err := models.GetCategorysName(o, category); err != nil {
		res.Msg = "資料庫錯誤！"
	} else {
		res.Status = 1
		res.Ext = category
	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) AddMonitoring() {

	title := this.GetString("title")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	userName := this.GetSession("uname")
	if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["ADD_IP_CATEGORY_MONITORING"], userName,"", title); err != nil {
		res.Msg = "資料庫錯誤1！"
		o.Rollback()
	} else {
		if id, err := models.AddCategory(o, title); err != nil {
			res.Msg = "資料庫錯誤2！"
			o.Rollback()
		} else {
			res.Status = 1
			res.Ext = id
			o.Commit()
		}
	}

	this.Data["json"] = res
	this.ServeJSON()

}

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
		o.Begin()

		userName := this.GetSession("uname")
		if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["ADD_IP_MONITORING"], userName,"", title); err != nil {
			res.Msg = "資料庫錯誤1！"
			o.Rollback()
		} else {
			categoryData := &models.TestPingCategory{}

			if err := models.GetCategory(o, categoryData, categoryId); err != nil {
				res.Msg = "資料庫錯誤2！"
				o.Rollback()
			} else {
				if _, err := models.AddIPMonitoring(o, categoryId, title, ip, t); err != nil {
					res.Msg = "資料庫錯誤3！"
					o.Rollback()
				} else {
					res.Status = 1
					o.Commit()
				}
			}
		}
	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) EditIPMonitoring() {

	id, _ := this.GetInt64("id")
	categoryId, _ := this.GetInt("category_id")
	title := this.GetString("title")
	ip := this.GetString("ip")
	t, _ := this.GetInt("type")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	data := models.EditIPMonitoringAfter{
		ID: id,
		Category: categoryId,
		Title: title,
		IP: ip,
		Type: t,
	}

	b, _ := json.Marshal(data)

	userName := this.GetSession("uname")
	if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["EDIT_IP_MONITORING"], userName,"", string(b)); err != nil {
		res.Msg = err.Error()
		o.Rollback()
	} else {
		categoryData := &models.TestPingCategory{}

		if err := models.GetCategory(o, categoryData, categoryId); err != nil {
			res.Msg = err.Error()
			o.Rollback()
		} else {
			if err := models.EditIPMonitoring(o, id, categoryId, t, title, ip); err != nil {
				res.Msg = err.Error()
				o.Rollback()
			} else {
				res.Status = 1
				o.Commit()
			}
		}
	}



	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) DelIPMonitoring() {

	id, _ := this.GetInt("id")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	userName := this.GetSession("uname")
	if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["DEL_IP_MONITORING"], userName,id, ""); err != nil {
		res.Msg = err.Error()
		o.Rollback()
	} else {
		if err := models.DelIPMonitoring(o, id); err != nil {
			res.Msg = "資料庫錯誤！"
		} else {
			res.Status = 1
			o.Commit()
		}
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

	var Ips []MonitoringItem

	o := orm.NewOrm()

	allData := []models.TestPingData{}
	if _, err := models.GetBaseAllTestPingData(o, &allData); err != nil {
		return
	}

	fmt.Printf("count:%d\n", len(Ips))
	for _, v := range allData {
		mystruct := MonitoringItem{
			CategoryId: v.CategoryId,
			Id: v.ItemId,
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

	itemId, _ := this.GetInt64("item_id")
	date := this.GetString("date")
	time := this.GetString("time")
	status, _ := this.GetInt("status")

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

		data := models.AddIPMonitoringResultAfter{
			categoryId, status, itemId, date, time, category, item, ip,
		}

		b, _ := json.Marshal(data)

		userName := this.GetSession("uname")
		if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["ADD_IP_MONITORING_RESULT"], userName,"", string(b)); err != nil {
			res.Msg = err.Error()
			o.Rollback()
		} else {
			if err := models.AddPingResult(o, categoryId, status, itemId, date, time, category, item, ip); err != nil {
				res.Msg = err.Error()
				o.Rollback()
			} else {
				res.Status = 1
				o.Commit()
			}
		}
	} else {
		res.Msg = "無法新增非當天報表資料！"
	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *IpMonitoringController) EditPingResult() {

	id, _ := strconv.Atoi(this.GetString("result_id"))
	status, _ := strconv.Atoi(this.GetString("status"))

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()
	o.Begin()

	result := &models.TestPingData{}

	if err := models.GetPingResult(o, result, id); err != nil {
		res.Msg = err.Error()
	} else {
		date := result.Date
		if isToday := CheckIsToday(date); isToday {
			beego.Debug(result.Status)
			before := models.EditIPMonitoringResultBefore{
				id, result.Status,
			}
			b, _ := json.Marshal(before)

			after := models.EditIPMonitoringResultAfter{
				id, status,
			}
			a, _ := json.Marshal(after)

			userName := this.GetSession("uname")
			if err := models.InsertOperationLog(o, conf.OPERATION_LOG_CODE["EDIT_IP_MONITORING_RESULT"], userName, string(b), string(a)); err != nil {
				res.Msg = err.Error()
				beego.Debug("1")
				o.Rollback()
			} else {
				if err := models.EditPingResult(o, id, status); err != nil {
					res.Msg = err.Error()
					o.Rollback()
				} else {
					res.Status = 1
					o.Commit()
				}
			}
		} else {
			// 非當日
			res.Msg = "無法修改非當天報表資料！"
		}
	}

	this.Data["json"] = res
	this.ServeJSON()

}