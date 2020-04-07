package main

import (
	"github.com/NavenduDuari/gomessenger/covid"
	"github.com/NavenduDuari/gomessenger/quote"
)

func main() {
	covid.SendCovidWs()
	quote.SendQuoteWs()
}
