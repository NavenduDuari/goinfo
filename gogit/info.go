package gogit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

var (
	reposBaseURL = "https://api.github.com/repos/"
	userBaseURL  = "https://api.github.com/users/"
)

func getCodeFrequencyURL(userName, repoName string) string {
	return reposBaseURL + userName + "/" + repoName + "/stats/code_frequency"
}

func getLanguagesURL(userName, repoName string) string {
	return reposBaseURL + userName + "/" + repoName + "/languages"
}

func getCommitURL(userName, repoName string) string {
	return reposBaseURL + userName + "/" + repoName + "/commits"
}
func getReposURL(userName string) string {
	return userBaseURL + userName + "/repos"
}

func getInfo(url string) {
	var bearer = "Bearer " + utils.GithubToken
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	utils.RawInfo <- body
	// return body
}

func getRepos(userName string) []utils.RepoStruct {
	ok := false
	var repos []utils.RepoStruct
	var rawRepos []byte
	fmt.Println("Getting repos from ", getReposURL(userName))
	go getInfo(getReposURL(userName))
	for {
		rawRepos, ok = <-utils.RawInfo
		if !ok {
			continue
		}
		break
	}
	json.Unmarshal(rawRepos, &repos)
	return repos
}
