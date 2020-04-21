package gogit

import (
	"encoding/json"
	"fmt"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

func getCodeFrequency(userName string) []utils.CodeFreqStruct {
	repos := getRepos(userName)
	fmt.Println("Repos => ", len(repos))
	var codeFreqs []utils.CodeFreqStruct
	go func() {
		for _, repo := range repos {
			codeFrequencyURL := getCodeFrequencyURL(userName, repo.Name)
			fmt.Println("LOC URL: ", codeFrequencyURL)
			go getInfo(codeFrequencyURL)

		}
	}()

	for i := 1; i <= len(repos); i++ {
		rawCodeFreq, ok := <-utils.RawInfo
		if !ok {
			continue
		}
		var codeFreq utils.CodeFreqStruct
		json.Unmarshal([]byte(rawCodeFreq), &codeFreq)
		codeFreqs = append(codeFreqs, codeFreq)
	}
	return codeFreqs
}

func GetLOC(userName string) int64 {
	var totalLOC int64
	codeFreqs := getCodeFrequency(userName)
	for _, codeFreq := range codeFreqs {
		for _, weeklyArr := range codeFreq {
			if len(weeklyArr) == 0 {
				continue
			}
			totalLOC = totalLOC + weeklyArr[1] + weeklyArr[2]
		}
	}
	// utils.GetLOC <- totalLOC
	return totalLOC
}
