package crypto

import "github.com/NavenduDuari/goinfo/crypto/utils"

func PrintPriceDown(currencySymbol string, price string, priceChangePercent string) string {
	return `Price: *` + currencySymbol + price + `* (` + priceChangePercent + `%) `
}

func PrintPriceUp(currencySymbol string, price string, priceChangePercent string) string {
	return `Price: *` + currencySymbol + price + `* (+` + priceChangePercent + `%) `
}

func PrintCoinInfo(id, name string) string {
	return `Coin: ` + name + `( *` + id + `* ) `
}

func PrintRank(rank string) string {
	return "Rank: *" + rank + "* \n"

}

func PrintCoinSuggestion() string {
	coinSuggestion := `*Use coin Id with --coin flag*` + " \n "
	for id, name := range utils.CoinDetails {
		coinSuggestion = coinSuggestion + "Id: *" + id + "*  Name: " + name + " \n "
	}
	return coinSuggestion
}

func PrintConvSuggestion() string {
	convSuggestion := `*Use coin Id with --conv flag*` + " \n "
	for id, details := range utils.CurrencyDetails {
		convSuggestion = convSuggestion + "Id: *" + id + "* Symbol: *" + details.Symbol + "*  Name: " + details.Name + " \n "
	}
	return convSuggestion
}
