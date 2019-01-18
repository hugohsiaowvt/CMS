package models

import (
	"github.com/astaxie/beego/orm"
)

type TestPingCategory struct {
	Id			int    `orm:"column(id)"`
	Title		string `orm:"column(title)"`
}

func (u *TestPingCategory) TableName() string {
	return "testpingcategory"
}

type TestPingItem struct {
	Id			int    `orm:"column(id)"`
	CategoryId	int    `orm:"column(category_id)"`
	Title		string `orm:"column(title)"`
	Ip			string `orm:"column(ip)"`
	Type		int    `orm:"column(type)"`
}

func (u *TestPingItem) TableName() string {
	return "testpingitem"
}

type TestPingResultData struct {
	Id			int    `orm:"column(result_id)"`
	ItemId		int    `orm:"column(item_id)"`
	Date		string	`orm:"column(date)"`
	Time		string	`orm:"column(time)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
	Ip			string	`orm:"column(ip)"`
	Status		int    `orm:"column(status)"`
}

func (u *TestPingResultData) TableName() string {
	return "testpingresult"
}

type TestPingData struct {
	CategoryId	int		`orm:"column(category_id)"`
	ItemId		int64	`orm:"column(item_id)"`
	Date		string	`orm:"column(date)"`
	Time		string	`orm:"column(time)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
	Ip			string	`orm:"column(ip)"`
	Type		int		`orm:"column(type)"`
	Status		int		`orm:"column(status)"`
}

type TestPingCategoryCount struct {
	Id			int		`orm:"column(id)"`
	Count		int		`orm:"column(count)"`
}

func GetCategory(o orm.Ormer, data *TestPingCategory, id int) error {
	return o.Raw("SELECT * FROM CMS.testpingcategory WHERE id = ?;", id).QueryRow(data)
}

func GetCategorysName(o orm.Ormer, data *[]TestPingCategory) (int64, error) {
	return o.Raw("SELECT * FROM CMS.testpingcategory;").QueryRows(data)
}

func GetBaseAllTestPingData(o orm.Ormer, data *[]TestPingData) (int64, error) {
	return o.Raw("SELECT ti.category_id, ti.id AS item_id, tc.title AS category, ti.title AS item, ti.ip as ip, ti.type as type FROM testpingcategory AS tc, testpingitem AS ti WHERE tc.id = ti.category_id ORDER BY tc.id, ti.id;").QueryRows(data)
}

func GetBaseCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) AS count  FROM testpingitem GROUP BY category_id;").QueryRows(data)
}

func GetTestPingResultByDate(o orm.Ormer, data *[]TestPingResultData, date1, date2 string) (int64, error) {
	return o.Raw("SELECT item_id, id AS result_id, time, status FROM testpingresult WHERE (date=? AND time >= '1830') OR (date=? AND time <= '0600') ORDER BY category_id, item_id;", date1, date2).QueryRows(data)
}

func GetPreviousPingData(o orm.Ormer, data *[]TestPingData, date1, date2 string) (int64, error) {
	return o.Raw("SELECT distinct item, category_id, item_id, category, ip FROM CMS.testpingresult WHERE (date = ? AND time >= '0600') OR (date = ? AND time < '0600') ORDER BY category_id, item_id;", date1, date2).QueryRows(data)
}

func GetPreviousCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount, date1, date2 string) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(distinct item) FROM CMS.testpingresult WHERE (date = ? AND time >= '0600') OR (date = ? AND time < '0600') GROUP BY category_id;", date1, date2).QueryRows(data)
}

func GetPingResult(o orm.Ormer, data *TestPingData, id int) error {
	return o.Raw("SELECT * FROM CMS.testpingresult WHERE id = ?;", id).QueryRow(data)
}

func InsertTestPingResult(o orm.Ormer, data TestPingData, date, time string, status int) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingresult` (`category_id`, `item_id`, `date`, `time`, `category`, `item`, `ip`, `status`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		data.CategoryId, data.ItemId, date, time, data.Category, data.Item, data.Ip, status).Exec()
	return err
}

func AddCategory(o orm.Ormer, title string) (int64, error) {
	res, err := o.Raw("INSERT INTO `CMS`.`testpingcategory` (`title`) VALUES (?);", title).Exec()
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func AddIPMonitoring(o orm.Ormer, c_id int, title, ip string, t int) (int64, error) {
	result, err := o.Raw("INSERT INTO `CMS`.`testpingitem` (`category_id`, `title`, `ip`, `type`) VALUES (?, ?, ?, ?);", c_id, title, ip, t).Exec()
	id, _ := result.LastInsertId()
	return id, err
}

func EditIPMonitoring(o orm.Ormer, id int64, c_id, t int, title, ip string) error {
	_, err := o.Raw("UPDATE `CMS`.`testpingitem` SET `category_id`=?, `title`=?, `ip`=?, `type`=? WHERE `id`=?;", c_id, title, ip, t, id).Exec()
	return err
}

func DelIPMonitoring(o orm.Ormer, id int) error {
	_, err := o.Raw("DELETE FROM `CMS`.`testpingitem` WHERE `id` = ?;", id).Exec()
	return err
}

func AddPingResult(o orm.Ormer, categoryId, status int, itemId int64, date, time, category, item, ip string) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingresult` (`category_id`, `item_id`, `date`, `time`, `category`, `item`, `ip`, `status`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		categoryId, itemId, date, time, category, item, ip, status).Exec()
	return err
}

func EditPingResult(o orm.Ormer, id, status int) error {
	_, err := o.Raw("UPDATE `CMS`.`testpingresult` SET `status` = ?  WHERE `id` = ?;", status, id).Exec()
	return err
}