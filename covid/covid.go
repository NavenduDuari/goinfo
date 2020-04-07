package covid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NavenduDuari/gomessenger/twilio"
	"github.com/NavenduDuari/gomessenger/utils"
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

func SendCovidWs() {
	covidObj := getCovidData()
	todayData := covidObj.Cases_time_series[len(covidObj.Cases_time_series)-1]
	msg := `Date: ` + todayData.Date + `
Total confirmed cases: ` + todayData.Totalconfirmed + "(+" + todayData.Dailyconfirmed + ")" + `
Total deceased: ` + todayData.Totaldeceased + "(+" + todayData.Dailydeceased + ")" + `
Total recovered: ` + todayData.Totalrecovered + "(+" + todayData.Dailyrecovered + ")" + `
Stay HOME, Stay SAFE` + `
	`

	for _, no := range utils.Contact {
		twilio.SendWhatsappMsg(no, msg)
	}
}
