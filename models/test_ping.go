package models

import (
	"github.com/astaxie/beego/orm"
)

type TestPingCategory struct {
	Id			int		`orm:"column(id)"`
	Category	string	`orm:"column(category)"`
	Status		int		`orm:"column(status)"`
}

func (u *TestPingCategory) TableName() string {
	return "testpingcategory"
}

type TestPingItem struct {
	Id			int		`orm:"column(id)"`
	CategoryId	int		`orm:"column(category_id)"`
	Item		string	`orm:"column(item)"`
	Status		int		`orm:"column(status)"`
}

func (u *TestPingItem) TableName() string {
	return "testpingitem"
}



type TestPingData struct {
	CategoryId	int		`orm:"column(category_id)"`
	ItemId		int		`orm:"column(item_id)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
	Ip			string	`orm:"column(ip)"`
}

type TestPingCategoryCount struct {
	Id			int		`orm:"column(id)"`
	Count		int		`orm:"column(count)"`
}

type TestPingResultData struct {
	ItemId		int		`orm:"column(item_id)"`
	ResultId	int		`orm:"column(result_id)"`
	Time		string	`orm:"column(time)"`
	Status		int		`orm:"column(status)"`
}

var TestPlanTime []string = []string {
	"PM6:30",
	"PM7:45",
	"PM8:15",
	"PM9:30",
	"PM10:45",
	"PM11:15",
	"AM00:30",
	"AM01:45",
	"AM02:15",
	"AM03:30",
	"AM04:45",
	"AM05:55",
}

func GetTestPlanTime() []string {
	return TestPlanTime;
}

func GetAllTestPingData(o orm.Ormer, data *[]TestPingData) (int64, error) {
	return o.Raw("SELECT ti.category_id, ti.id AS item_id, tc.title AS category, ti.title AS item, ti.ip as ip FROM testpingcategory AS tc, testpingitem AS ti WHERE tc.id = ti.category_id ORDER BY ti.id;").QueryRows(data)
}

func GetCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) AS count  FROM testpingitem GROUP BY category_id;").QueryRows(data)
}

func GetTestPingResultByDate(o orm.Ormer, data *[]TestPingResultData, date1, date2 string) (int64, error) {
	return o.Raw("SELECT item_id, id AS result_id, time, status FROM testpingresult WHERE (date=? AND time >= '1830') OR (date=? AND time <= '0600');", date1, date2).QueryRows(data)
}

func InsertTestPingResult(o orm.Ormer, data TestPingData, date, time string, status int) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingresult` (`item_id`, `date`, `time`, `category`, `item`, `ip`, `status`) VALUES (?, ?, ?, ?, ?, ?, ?);",
		data.ItemId, date, time, data.Category, data.Item, data.Ip, status).Exec()
	return err
}