package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/NavenduDuari/goinfo/utils"
)

type coinData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	OneDay oneDay `json:"1D"`
}

type oneDay struct {
	PriceChange string `json:"price_change"`
}

var (
	baseUrl        = "https://api.nomics.com/v1/currencies/ticker?key="
	currencySymbol = CurrencyDetails["INR"].Symbol
)

func getPrice(w http.ResponseWriter, coin, conv string) {
	var (
		ids     = "&ids=BTC,ETH,XRP"
		convert = "&convert=INR"
		key     = utils.NomicsApiKey
	)
	if coin != "" {
		ids = "&ids=" + coin
	}
	if conv != "" {
		convert = "&convert=" + conv
		currencySymbol = CurrencyDetails[conv].Symbol
	}
	finalUrl := baseUrl + key + ids + "&interval=1d" + convert

	res, err := http.Get(finalUrl)
	if err != nil {
		fmt.Println("Unable to get price")
		return
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	var coinDataArrObj []coinData
	json.Unmarshal(responseData, &coinDataArrObj)

	showPrice(w, coinDataArrObj)
}

func showPrice(w http.ResponseWriter, coinDataArrObj []coinData) {
	content := ""
	for _, coin := range coinDataArrObj {
		priceChange, _ := strconv.ParseFloat(coin.OneDay.PriceChange, 64)
		price, _ := strconv.ParseFloat(coin.Price, 64)
		priceChangePercent := fmt.Sprintf("%.2f", priceChange/(priceChange+price)*100)
		content = content + PrintCoinInfo(coin.Id, coin.Name)
		if priceChange < 0 {
			content = content + PrintPriceDown(currencySymbol, coin.Price, priceChangePercent)
		} else {
			content = content + PrintPriceUp(currencySymbol, coin.Price, priceChangePercent)
		}
		content = content + PrintRank(coin.Rank)
	}
	io.WriteString(w, content)
}

func getSuggestion(w http.ResponseWriter) {
	content := PrintCoinSuggestion() + "\n" + PrintConvSuggestion()
	io.WriteString(w, content)

}
func getHelp(w http.ResponseWriter) {
	content := `*crypto* gives prices of crypto-currencies.` + " \n " + `
	commands available:
	*--coin*		//to specify coin
	*--conv*		//to specify conversion
	*--help*		//to get help
	*--suggest*	//to get suggestion

	Example:
	*crypto* 		//gives price default coins in default conversion
	*crypto --coin=BTC,LTC,BNB --conv=EUR*		//gives price of LTC in EUR`

	io.WriteString(w, content)
}

func Check(w http.ResponseWriter, args map[string]string, isCmdValid bool) {
	if args["--suggest"] == "true" {
		getSuggestion(w)
	} else if args["--help"] == "true" || isCmdValid == false {
		getHelp(w)
	} else {
		getPrice(w, args["--coin="], args["--conv="])
	}
}
