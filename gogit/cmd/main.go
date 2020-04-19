package main

import (
	"fmt"

	"github.com/NavenduDuari/goinfo/gogit"
)

func main() {
	fmt.Println("Total LOC => ", gogit.GetLOC())
	fmt.Println("Language => ", gogit.GetLanguagePercentage())
	fmt.Println("Commit => ", gogit.GetCommitCount())
}
