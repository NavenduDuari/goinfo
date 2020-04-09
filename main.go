package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/NavenduDuari/goinfo/utils"

	"github.com/NavenduDuari/goinfo/covid"
	"github.com/NavenduDuari/goinfo/crypto"
	"github.com/NavenduDuari/goinfo/quote"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		// resByte, _ := ioutil.ReadAll(r.Body)
		// responseMap := utils.DecodeResponse(string(resByte))
		// body, _ := url.QueryUnescape(responseMap["Body"])

		// recognizeCommandAndCall(w, body)

		io.WriteString(w, "*Under Maintenance. Please try after sometime. Sorry!*")
	}

	http.HandleFunc("/endpoint", h)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	fmt.Println(port)
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

	other := ""
	switch c.cmd {
	case "gocoin":
		if c.args != nil {
			other = "true"
		} else {
			for arg := range c.args {
				for _, validArg := range utils.CryptoArgs {
					if validArg == arg {
						other = ""
						break
					}
				}
			}
		}
		crypto.Check(w, c.args["--coin="], c.args["--conv="], c.args["--suggest"], c.args["--help"], other)
	case "covid":
		if c.args != nil {
			other = "true"
		} else {
			for arg := range c.args {
				for _, validArg := range utils.CovidArgs {
					if validArg == arg {
						other = ""
						break
					}
				}
			}
		}
		covid.SendCovidWs(w, c.args["--state="], c.args["--suggest"], c.args["--help"], other)
	case "quote":
		if c.args != nil {
			other = "true"
		} else {
			for arg := range c.args {
				for _, validArg := range utils.QuoteArgs {
					if validArg == arg {
						other = ""
						break
					}
				}
			}
		}
		quote.SendQuoteWs(w, c.args["--cat="], c.args["--suggest"], c.args["--help"], other)
	default:
		content := `Try:
			gocoin --help
			covid --help
			quote --help`
		io.WriteString(w, content)
	}
}
