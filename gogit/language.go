package gogit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

var (
	totalByteCount            float64
	languageWithByteStructArr []*utils.LanguageWithByteStruct
)

func getLanguages(w http.ResponseWriter) []utils.LanguageWithByteMap {
	fmt.Println("within getLanguages")
	io.WriteString(w, "within getLanguages")

	var languages []utils.LanguageWithByteMap
	repos := getRepos(w)
	for _, repo := range repos {
		var lang utils.LanguageWithByteMap
		languagesURL := getLanguagesURL(repo.Name)
		rawLang := getInfo(w, languagesURL)
		json.Unmarshal([]byte(rawLang), &lang)
		languages = append(languages, lang)
	}
	return languages
}

func calculateLanguageByte(w http.ResponseWriter) []*utils.LanguageWithByteStruct {
	fmt.Println("within calculateLanguageByte")
	io.WriteString(w, "within calculateLanguageByte")

	languageWithByteMapArr := getLanguages(w)
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

	return languageWithByteStructArr
}

func GetLanguagePercentage(w http.ResponseWriter) []utils.LanguageWithPercentageStruct {
	fmt.Println("within GetLanguagePercentage")
	io.WriteString(w, "within GetLanguagePercentage")

	var languageWithPercentageStructArr []utils.LanguageWithPercentageStruct

	languageWithByteStructArr := calculateLanguageByte(w)
	for _, languageWithByteStruct := range languageWithByteStructArr {
		languageWithPercentageStructArr = append(languageWithPercentageStructArr,
			utils.LanguageWithPercentageStruct{
				Language:   languageWithByteStruct.Language,
				Percentage: (languageWithByteStruct.ByteCount / totalByteCount * 100),
			})
	}
	return languageWithPercentageStructArr
}
