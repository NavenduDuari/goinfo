package gogit

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/NavenduDuari/goinfo/gogit/utils"
)

func getCommit(w http.ResponseWriter) []utils.CommitStruct {
	io.WriteString(w, "within getCommit")

	var commitStructArr []utils.CommitStruct
	repos := getRepos(w)
	for _, repo := range repos {
		var commitStruct utils.CommitStruct
		commitURL := getCommitURL(repo.Name)
		rawCommitInfo := getInfo(w, commitURL)
		json.Unmarshal([]byte(rawCommitInfo), &commitStruct)
		commitStructArr = append(commitStructArr, commitStruct)
	}
	return commitStructArr
}

func GetCommitCount(w http.ResponseWriter) int {
	io.WriteString(w, "within GetCommitCount")

	var totalCommit int
	commitStructArr := getCommit(w)
	for _, commitStruct := range commitStructArr {
		totalCommit += len(commitStruct)
	}
	return totalCommit
}
