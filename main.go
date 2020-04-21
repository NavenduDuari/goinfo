package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/NavenduDuari/goinfo/gogit"
	"github.com/NavenduDuari/goinfo/utils"

	"github.com/NavenduDuari/goinfo/covid"
	covidUtils "github.com/NavenduDuari/goinfo/covid/utils"
	"github.com/NavenduDuari/goinfo/crypto"
	cryptoUtils "github.com/NavenduDuari/goinfo/crypto/utils"
	gogitUtils "github.com/NavenduDuari/goinfo/gogit/utils"
	"github.com/NavenduDuari/goinfo/quote"
	quoteUtils "github.com/NavenduDuari/goinfo/quote/utils"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		resByte, _ := ioutil.ReadAll(r.Body)
		responseMap := utils.DecodeResponse(string(resByte))
		body, _ := url.QueryUnescape(responseMap["Body"])

		recognizeCommandAndCall(w, body)

		// io.WriteString(w, "*Under Maintenance. Please try after sometime. Sorry!*")
	}

	http.HandleFunc("/endpoint", h)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	log.Fatal(http.ListenAndServe(port, nil))
}

type command struct {
	cmd  string
	args map[string]string
}

func recognizeCommandAndCall(w http.ResponseWriter, cmdStr string) {
	var c command
	c.args = make(map[string]string)
	words := strings.Split(cmdStr, " ")
	c.cmd = words[0]
	for i := 1; i < len(words); i++ {
		temp := strings.SplitAfter(words[i], "=")
		if len(temp) < 2 {
			c.args[temp[0]] = "true"
		} else {
			c.args[temp[0]] = temp[1]
		}
	}

	switch c.cmd {
	case "crypto":
		fmt.Println("crypto")
		if cryptoUtils.IsCmdValid(c.args) {
			crypto.Check(w, c.args, true)
		} else {
			crypto.Check(w, c.args, false)
		}
	case "covid":
		fmt.Println("covid")
		if covidUtils.IsCmdValid(c.args) {
			covid.SendCovidWs(w, c.args, true)
		} else {
			covid.SendCovidWs(w, c.args, false)
		}
	case "quote":
		fmt.Println("quote")
		if quoteUtils.IsCmdValid(c.args) {
			quote.SendQuoteWs(w, c.args, true)
		} else {
			quote.SendQuoteWs(w, c.args, false)
		}
	case "gogit":
		fmt.Println("gogit")
		if gogitUtils.IsCmdValid(c.args) {
			gogit.ServeGogit(w, c.args, true)
		} else {
			gogit.ServeGogit(w, c.args, false)
		}
	default:
		fmt.Println("default")
		content := `Welcome to *GOINFO*
		Try:
		*crypto --help*
		*covid --help*
		*quote --help*
		*gogit --help*`
		io.WriteString(w, content)
	}
}
