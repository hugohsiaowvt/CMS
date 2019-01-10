package models

type Page struct {
	Tag	, LinkPage , IconPath		string
}

var (
	PageChilds []Page = []Page {
		Page{
			Tag:"問題回報",
			LinkPage:"welcome.tpl",
			IconPath:"/static/img/report.png",
		},
		Page{
			Tag:"扛打監聽",
			LinkPage:"test_ping.tpl",
			IconPath:"/static/img/monitor.png",
		},
		Page{
			Tag:"排程設定",
			LinkPage:"ip_monitoring.tpl",
			IconPath:"/static/img/schedule.png",
		},
	}
)

func GeneratePages() []Page {
	return PageChilds;
}