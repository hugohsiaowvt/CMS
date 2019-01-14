package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type ResponseStatus struct {
	Status int
	Ext	   interface{}
	Msg	   string
}

func CheckIsToday(date string) bool {
	// 判斷是否為今天報表
	// 1.如果日期是昨天而且今天的時間大於0600就不是今日報表
	// 2.如果日期差兩天以上
	t := time.Now().Local()
	subDay, _ := time.ParseDuration("-24h")
	tmpDay := t.Add(subDay)
	yesterDay := tmpDay.Format("2006-01-02")
	nowTime := t.Format("1504")

	beego.Debug(date, yesterDay, nowTime)

	isToday := true
	if date == yesterDay && nowTime > "0600" || date < yesterDay {
		isToday = false
	}

	return isToday
}