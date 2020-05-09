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
	stateInfo := ""
	for _, stateData := range covidObj.Statewise {
		if stateData.Statecode == args["--state="] {
			todayData = stateData
			stateFound = true
			stateInfo = `State: ` + todayData.State + "(" + todayData.Statecode + ")"
			break
		}
	}
	todayData = covidObj.Statewise[0] //TOTAL WARN: dependent on struct

	msg = `Last Updated: ` + todayData.Lastupdatedtime + `
		` + stateInfo + `
Total confirmed cases: ` + todayData.Confirmed + "(+" + todayData.Deltaconfirmed + ")" + `
Total deceased: ` + todayData.Deaths + "(+" + todayData.Deltadeaths + ")" + `
Total recovered: ` + todayData.Recovered + "(+" + todayData.Deltarecovered + ")" + `
Stay HOME, Stay SAFE` + `
		`

	io.WriteString(w, msg)

	if !stateFound {
		getSuggestion(w)
	}
}

func getSuggestion(w http.ResponseWriter) {
	content := `-------------------------------------------------------------
	TRY *covid --state=Id*
	`
	for stateId, stateName := range utils.States {
		content = content + `
		Name: ` + stateName + ` Id: *` + stateId + `*
`
	}
	io.WriteString(w, content)
}
