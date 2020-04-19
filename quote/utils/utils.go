package utils

var QuoteArgs = []string{"--cat=", "--suggest", "--help"}

var QuoteCategory = []string{
	"inspire",
	"management",
	"life",
	"love",
	"art",
	"students",
}

func IsCmdValid(argsMap map[string]string) bool {
	if len(argsMap) == 0 {
		return true
	}

	for _, validArg := range QuoteArgs {
		for arg := range argsMap {
			if validArg == arg {
				return true
			}
		}
	}
	return false
}
