package quote

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/NavenduDuari/goinfo/utils"
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

func SendQuoteWs(w http.ResponseWriter, category, suggest, help interface{}) {
	if help != nil {
		content := `*quote* gives you great quotes.` + " \n " + `
		commands available:
		--cat		//to specify category
		--help		//to get help
		--suggest	//to get suggestion
	
		Example:
		quote		//gives random quote
		quote --cat=inspire		//gives quote of inspire category`

		io.WriteString(w, content)
	} else if suggest != nil {
		content := `*Available categories are:*
		`
		for _, category := range utils.QuoteCategory {
			content = content + category + `
			`
		}
		io.WriteString(w, content)
	} else if category != nil {
		response := getQuote(category.(string))
		quote := response.Contents.Quotes[0]
		msg := quote.Quote + `
-- ` + "*" + quote.Author + "*"

		io.WriteString(w, msg)
	} else {
		rand.Seed(time.Now().UnixNano())
		cat := utils.QuoteCategory[rand.Intn(len(utils.QuoteCategory))]
		response := getQuote(cat)
		quote := response.Contents.Quotes[0]
		msg := quote.Quote + `
-- ` + "*" + quote.Author + "*"

		io.WriteString(w, msg)
	}
	// for _, no := range utils.Contact {
	// 	twilio.SendWhatsappMsg(no, msg)
	// }
}
