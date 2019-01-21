package conf

var (

	PING_STATUS_OK = 1
	PING_STATUS_NOTOK = -1

	PING_TIME = [] string{
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

	TEST_PLAN_TIME = []string {
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

	OPERATION_LOG_CODE = map[string]string {

		"ADD_IP_CATEGORY_MONITORING":	"101",
		"ADD_IP_MONITORING":			"102",
		"ADD_IP_MONITORING_RESULT":		"103",

		"EDIT_IP_CATEGORY_MONITORING":	"201",
		"EDIT_IP_MONITORING":			"202",
		"EDIT_IP_MONITORING_RESULT":	"203",

		"DEL_IP_CATEGORY_MONITORING":	"301",
		"DEL_IP_MONITORING":			"302",

	}
)