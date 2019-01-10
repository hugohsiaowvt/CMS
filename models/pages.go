package models

type Page struct {
	Tag	, LinkPage		string
}

var (
	PageChilds []Page = []Page {
		Page{
			Tag:"Page1",
			LinkPage:"testPage1.tpl",
		},

		Page{
			Tag:"Page2",
			LinkPage:"testPage2.tpl",
		},

		Page{
			Tag:"Page3",
			LinkPage:"welcome.tpl",
		},
		Page{
			Tag:"問題回報",
			LinkPage:"welcome.tpl",
		},
		Page{
			Tag:"扛打監聽",
			LinkPage:"test_ping.tpl",
		},
		Page{
			Tag:"排程設定",
			LinkPage:"ip_monitoring.tpl",
		},
	}
)

func GeneratePages() []Page {
	return PageChilds;
}