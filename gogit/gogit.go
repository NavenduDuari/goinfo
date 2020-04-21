package gogit

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strconv"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

func getHelp(w http.ResponseWriter) {
	content := `*gogit* gives basic insights of Github profile.` + " \n " + `
		commands available:
		*--name*		//to specify username
		*--help*		//to get help
	
		Example:
		*gogit*		//gives my github insights
		*gogit --name=subhash1991*		//gives insights from github profile of 'subhash1991'`

	io.WriteString(w, content)
}

func getInsight(userName string) string {
	emojiFingure := html.UnescapeString("&#" + strconv.Itoa(128073) + ";")
	emojiAvatar := html.UnescapeString("&#" + strconv.Itoa(128100) + ";")
	var langResp = emojiFingure + ` Languages(%) => `
	langauges := GetLanguagePercentage(userName)
	for _, lang := range langauges {
		langResp += fmt.Sprintf("\n\t\t\t\t%s -- *%.2f*", lang.Language, lang.Percentage)
	}
	locResp := fmt.Sprintf("%s Total LOC => *%d* ", emojiFingure, GetLOC(userName))
	commitResp := fmt.Sprintf("%s Total Commits => *%d* ", emojiFingure, GetCommitCount(userName))
	response := `
	` + emojiAvatar + ` *` + userName + `*
	-----------------------------------------------
	` + locResp + `
	` + commitResp + `
	` + langResp

	return response
}

func ServeGogit(w http.ResponseWriter, args map[string]string, isCmdValid bool) {
	if args["--help"] != "" || isCmdValid == false {
		getHelp(w)
	} else if args["--name="] != "" {
		response := getInsight(args["--name="])
		io.WriteString(w, response)
	} else {
		response := getInsight(utils.DefaultUserName)
		io.WriteString(w, response)
	}
}
