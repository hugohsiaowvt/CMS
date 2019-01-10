package controllers

import (
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

	date1 := this.GetString("date1")
	date2 := this.GetString("date2")

	if date1 == "" || date2 == "" {
		t := time.Now().Local()
		date1 = t.Format("20060102")
		nextDay, _ := time.ParseDuration("24h")
		t = t.Add(nextDay)
		date2 = t.Format("20060102")
		beego.Debug(date1, date2)
	}

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

	beego.Debug(resultData)

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

