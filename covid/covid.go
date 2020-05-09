package covid

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/NavenduDuari/goinfo/covid/utils"
)

func getCovidData() covidStruct {
	url := "https://api.covid19india.org/data.json"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var covidObj covidStruct
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(responseData, &covidObj)

	return covidObj
}
func getHelp(w http.ResponseWriter) {
	content := `*covid* gives covid data.` + " \n " + `
	commands available:
	*--state*		//to specify statae
	*--help*		//to get help
	*--suggest*	//to get suggestion

	Example:
	*covid*		//gives data of India
	*covid --state=WB*		//gives data of West Bengal`

	io.WriteString(w, content)
}
func SendCovidWs(w http.ResponseWriter, args map[string]string, isCmdValid bool) {
	var msg string
	var todayData statewiseCases
	covidObj := getCovidData()
	if args["--suggest"] != "" {
		getSuggestion(w)
	} else if args["--help"] != "" || isCmdValid == false {
		getHelp(w)
	}

	stateFound := false
	for _, stateData := range covidObj.Statewise {
		if stateData.Statecode == args["--state="] {
			todayData = stateData
			stateFound = true
		}
	}
	todayData = covidObj.Statewise[0] //TOTAL WARN: dependent on struct
	if !stateFound {
		getSuggestion(w)
	}
	msg = `Last Updated: ` + todayData.Lastupdatedtime + `
State: ` + todayData.State + "(" + todayData.Statecode + ")" + `
Total confirmed cases: ` + todayData.Confirmed + "(+" + todayData.Deltaconfirmed + ")" + `
Total deceased: ` + todayData.Deaths + "(+" + todayData.Deltadeaths + ")" + `
Total recovered: ` + todayData.Recovered + "(+" + todayData.Deltarecovered + ")" + `
Stay HOME, Stay SAFE` + `
		`

	// 	} else if args["--state="] != "" {
	// 		var stateFound bool
	// 		for _, stateData := range covidObj.Statewise {
	// 			if stateData.Statecode == args["--state="] {
	// 				todayData = stateData
	// 				stateFound = true
	// 			}
	// 		}
	// 		if !stateFound {
	// 			getSuggestion(w)
	// 			return
	// 		}
	// 		msg = `Last Updated: ` + todayData.Lastupdatedtime + `
	// State: ` + todayData.State + "(" + todayData.Statecode + ")" + `
	// Total confirmed cases: ` + todayData.Confirmed + "(+" + todayData.Deltaconfirmed + ")" + `
	// Total deceased: ` + todayData.Deaths + "(+" + todayData.Deltadeaths + ")" + `
	// Total recovered: ` + todayData.Recovered + "(+" + todayData.Deltarecovered + ")" + `
	// Stay HOME, Stay SAFE` + `
	// 		`
	// 	} else {
	// 		todayData := covidObj.Cases_time_series[len(covidObj.Cases_time_series)-1]
	// 		msg = `Last Updated: ` + todayData.Date + `
	// Total confirmed cases: ` + todayData.Totalconfirmed + "(+" + todayData.Dailyconfirmed + ")" + `
	// Total deceased: ` + todayData.Totaldeceased + "(+" + todayData.Dailydeceased + ")" + `
	// Total recovered: ` + todayData.Totalrecovered + "(+" + todayData.Dailyrecovered + ")" + `
	// Stay HOME, Stay SAFE` + `
	// 	`
	// 	}

	io.WriteString(w, msg)
}

func getSuggestion(w http.ResponseWriter) {
	content := ""
	for stateId, stateName := range utils.States {
		content = content + `
		Name: ` + stateName + ` Id: *` + stateId + `*
`
	}
	io.WriteString(w, content)
}
