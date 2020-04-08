package covid

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/NavenduDuari/goinfo/utils"
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

func SendCovidWs(w http.ResponseWriter, state, suggest, help string) {
	var msg string
	var todayData statewiseCases
	covidObj := getCovidData()
	if suggest != "" {
		getSuggestion(w)
	} else if help != "" {
		content := `*covid* gives covid data.` + " \n " + `
		commands available:
		--state		//to specify statae
		--help		//to get help
		--suggest	//to get suggestion
	
		Example:
		covid		//gives data of India
		covid --state=WB		//gives data of West Bengal`

		io.WriteString(w, content)
	} else if state != "" {
		for _, stateData := range covidObj.Statewise {
			if stateData.Statecode == state {
				todayData = stateData
			}
		}
		msg = `Last Updated: ` + todayData.Lastupdatedtime + `
State: ` + todayData.State + "(" + todayData.Statecode + ")" + `
Total confirmed cases: ` + todayData.Confirmed + "(+" + todayData.Deltaconfirmed + ")" + `
Total deceased: ` + todayData.Deaths + "(+" + todayData.Deltadeaths + ")" + `
Total recovered: ` + todayData.Recovered + "(+" + todayData.Deltarecovered + ")" + `
Stay HOME, Stay SAFE` + `
		`
	} else {
		todayData := covidObj.Cases_time_series[len(covidObj.Cases_time_series)-1]
		msg = `Last Updated: ` + todayData.Date + `
Total confirmed cases: ` + todayData.Totalconfirmed + "(+" + todayData.Dailyconfirmed + ")" + `
Total deceased: ` + todayData.Totaldeceased + "(+" + todayData.Dailydeceased + ")" + `
Total recovered: ` + todayData.Totalrecovered + "(+" + todayData.Dailyrecovered + ")" + `
Stay HOME, Stay SAFE` + `
	`
	}

	io.WriteString(w, msg)
	// for _, no := range utils.Contact {
	// 	twilio.SendWhatsappMsg(no, msg)
	// }
}

func getSuggestion(w http.ResponseWriter) {
	content := ""
	for stateId, stateName := range utils.States {
		content = content + `
Id: *` + stateId + `* Name: ` + stateName + `
`
	}
	io.WriteString(w, content)
}
