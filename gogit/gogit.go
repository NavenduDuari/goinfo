package gogit

import (
	"fmt"
	"io"
	"net/http"
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

func getInsight(w http.ResponseWriter, name string) string {
	fmt.Println("within getInsight")
	io.WriteString(w, "within getInsight")

	var langResp = ""
	langauges := GetLanguagePercentage(w)
	for _, lang := range langauges {
		langResp += fmt.Sprintf("%s -- %.2f \n", lang.Language, lang.Percentage)
	}
	locResp := fmt.Sprintf("Total LOC => %d", GetLOC(w))
	commitResp := fmt.Sprintf("Total Commits => %d", GetCommitCount(w))
	response := locResp + "\n" + commitResp + "\n" + langResp

	// fmt.Println(response)
	return response
}

func ServeGogit(w http.ResponseWriter, args map[string]string, isCmdValid bool) {
	fmt.Println("within ServeGogit")
	io.WriteString(w, "within ServeGogit")

	if args["--help"] != "" || isCmdValid == false {
		getHelp(w)
	} else if args["--name="] != "" {
		io.WriteString(w, "with name")
		io.WriteString(w, "with name")
		fmt.Println("within else if args == --name")
		// response := getInsight(w, args["--name"])
		// io.WriteString(w, response)
	} else {
		io.WriteString(w, "default name")
		io.WriteString(w, "default name")
		// response := getInsight(w, utils.DefaultUserName)
		// io.WriteString(w, response)

	}
}
