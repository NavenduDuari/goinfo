package gogit

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

func getCodeFrequency(w http.ResponseWriter) []utils.CodeFreqStruct {
	io.WriteString(w, "within getCodeFrequency")

	repos := getRepos(w)
	var codeFreqs []utils.CodeFreqStruct
	for _, repo := range repos {
		var codeFreq utils.CodeFreqStruct
		codeFrequencyURL := getCodeFrequencyURL(repo.Name)
		rawCodeFreq := getInfo(w, codeFrequencyURL)
		json.Unmarshal([]byte(rawCodeFreq), &codeFreq)
		codeFreqs = append(codeFreqs, codeFreq)
	}
	return codeFreqs
}

func GetLOC(w http.ResponseWriter) int64 {
	io.WriteString(w, "within GetLOC")

	var totalLOC int64
	codeFreqs := getCodeFrequency(w)
	for _, codeFreq := range codeFreqs {
		for _, weeklyArr := range codeFreq {
			totalLOC = totalLOC + weeklyArr[1] + weeklyArr[2]
		}
	}
	return totalLOC
}
