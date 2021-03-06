package covid

type covidStruct struct {
	Cases_time_series []cases          `json:"cases_time_series"`
	Statewise         []statewiseCases `json:"statewise"`
	// Tested            []testedCases    `json:tested`
}

type testedCases struct {
}

type statewiseCases struct {
	Active          string `json:"active"`
	Confirmed       string `json:"confirmed"`
	Deaths          string `json:"deaths"`
	Deltaconfirmed  string `json:"deltaconfirmed"`
	Deltadeaths     string `json:"deltadeaths"`
	Deltarecovered  string `json:"deltarecovered"`
	Lastupdatedtime string `json:"lastupdatedtime"`
	Recovered       string `json:"recovered"`
	State           string `json:"state"`
	Statecode       string `json:"statecode"`
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
