package gogit

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

var (
// totalByteCount float64
)

func getLanguages(userName string) []utils.LanguageWithByteMap {
	var languages []utils.LanguageWithByteMap
	repos := getRepos(userName)
	go func() {
		for _, repo := range repos {
			languagesURL := getLanguagesURL(userName, repo.Name)
			fmt.Println("Lang url : ", languagesURL)
			go getInfo(languagesURL)
		}
	}()

	for i := 1; i <= len(repos); i++ {
		rawLang, ok := <-utils.RawInfo
		if !ok {
			continue
		}
		var lang utils.LanguageWithByteMap
		json.Unmarshal([]byte(rawLang), &lang)
		languages = append(languages, lang)
	}

	return languages
}

func calculateLanguageByte(userName string) ([]*utils.LanguageWithByteStruct, float64) {
	var totalByteCount float64
	var languageWithByteStructArr []*utils.LanguageWithByteStruct
	languageWithByteMapArr := getLanguages(userName)
	for _, languageWithByteMap := range languageWithByteMapArr {
		for lang, byteCount := range languageWithByteMap {
			isUpdated := false
			if len(languageWithByteStructArr) == 0 {
				languageWithByteStructArr = append(languageWithByteStructArr,
					&utils.LanguageWithByteStruct{
						Language:  lang,
						ByteCount: byteCount,
					})
				totalByteCount += byteCount
				isUpdated = true
			} else {
				for _, languageWithByteStruct := range languageWithByteStructArr {
					if lang == languageWithByteStruct.Language {
						languageWithByteStruct.ByteCount += byteCount
						totalByteCount += byteCount
						isUpdated = true
					}
				}
				if !isUpdated {
					languageWithByteStructArr = append(languageWithByteStructArr,
						&utils.LanguageWithByteStruct{
							Language:  lang,
							ByteCount: byteCount,
						})
					totalByteCount += byteCount
				}
			}
		}
	}

	return languageWithByteStructArr, totalByteCount
}

func GetLanguagePercentage(userName string) []utils.LanguageWithPercentageStruct {
	var languageWithPercentageStructArr []utils.LanguageWithPercentageStruct

	languageWithByteStructArr, totalByteCount := calculateLanguageByte(userName)
	for _, languageWithByteStruct := range languageWithByteStructArr {
		languageWithPercentageStructArr = append(languageWithPercentageStructArr,
			utils.LanguageWithPercentageStruct{
				Language:   languageWithByteStruct.Language,
				Percentage: (languageWithByteStruct.ByteCount / totalByteCount * 100),
			})
	}
	sort.SliceStable(languageWithPercentageStructArr, func(i, j int) bool {
		return languageWithPercentageStructArr[i].Percentage > languageWithPercentageStructArr[j].Percentage
	})
	// utils.GetLang <- languageWithPercentageStructArr
	return languageWithPercentageStructArr
}
