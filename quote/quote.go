package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/NavenduDuari/gomessenger/twilio"
	"github.com/NavenduDuari/gomessenger/utils"
)

func getQuote(category string) responseStruct {
	url := "http://quotes.rest/qod.json?category=" + category

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var responseObj responseStruct
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(responseData, &responseObj)

	return responseObj
}

func SendQuoteWs() {
	rand.Seed(time.Now().UnixNano())
	category := utils.QuoteCategory[rand.Intn(len(utils.QuoteCategory))]
	response := getQuote(category)
	quote := response.Contents.Quotes[0]
	msg := quote.Quote + `
-- ` + "*" + quote.Author + "*"

	for _, no := range utils.Contact {
		twilio.SendWhatsappMsg(no, msg)
	}
}
