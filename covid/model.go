package covid

type covidStruct struct {
	Cases_time_series []cases `json:"cases_time_series"`
	// Statewise         []statewiseCases `json:"statewise"`
	// Tested            []testedCases    `json:tested`
}

type testedCases struct {
}

type statewiseCases struct {
	Active string `json:"active"`
	//   "confirmed": "5337",
	//   "deaths": "154",
	//   "deltaconfirmed": "559",
	//   "deltadeaths": "20",
	//   "deltarecovered": "51",
	//   "lastupdatedtime": "07/04/2020 22:44:25",
	//   "recovered": "445",
	//   "state": "Total",
	//   "statecode": "TT"
}
type cases struct {
	Dailyconfirmed string `json:"dailyconfirmed"`
	Dailydeceased  string `json:"dailydeceased"`
	Dailyrecovered string `json:"dailyrecovered"`
	Date           string `json:"date"`
	Totalconfirmed string `json:"totalconfirmed"`
	Totaldeceased  string `json:"totaldeceased"`
	Totalrecovered string `json:"totalrecovered"`
}
