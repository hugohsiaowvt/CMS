package controllers

import (
	"CMS/conf"
	"CMS/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
)

type DashBoardController struct {
	beego.Controller
}


func (this *DashBoardController) TestPing() {

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

	this.TplName = "test_ping.tpl"
}

type ShowPageController struct {
	beego.Controller
}

func (c *DashBoardController) Get() {
	Pages:= models.GeneratePages();
	c.Data["Pages"]=Pages
	c.TplName = "dashboard.tpl"
	c.Data["testHello"]="test";
	//內容頁面
	/*c.Layout = "dashboard.tpl"
	page := c.GetString("page")
	if page == "" {
		c.TplName = "testPage1.tpl"
	} else {
		fmt.Printf("\nkey:%s\n", page)
		c.TplName = page
	}*/

}

func (this *ShowPageController) ShowPage() {
	//this.TplName = "dashboard.tpl"
	//內容頁面
	//this.Layout = "dashboard.tpl"
	page := this.GetString("page")
	if page == "" {
		this.TplName = "testPage1.tpl"
	} else {
		fmt.Printf("\nkey:%s\n", page)
		this.TplName = page
	}
}

