package utils

type RepoStruct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type LanguageWithByteMap map[string]float64

type LanguageWithByteStruct struct {
	Language  string
	ByteCount float64
}

type LanguageWithPercentageStruct struct {
	Language   string
	Percentage float64
}

type CodeFreqStruct [][]int64

type CommitStruct []CommitBodystruct

type CommitBodystruct struct {
	SHA string `json:"sha"`
}

var DefaultUserName = "NavenduDuari"
var GogitArgs = []string{"--name=", "--help"}

func IsCmdValid(argsMap map[string]string) bool {
	if len(argsMap) == 0 {
		return true
	}
	for _, validArg := range GogitArgs {
		for arg := range argsMap {
			if validArg == arg {
				return true
			}
		}
	}
	return false
}
