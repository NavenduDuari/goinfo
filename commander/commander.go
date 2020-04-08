package commander

import (
	"fmt"
	"os/exec"
)

func Execute(cmd string) string {
	output, err := exec.Command(cmd).Output()
	if err != nil {
		output = []byte(err.Error())
		fmt.Println(err)
	}
	return string(output)
}
