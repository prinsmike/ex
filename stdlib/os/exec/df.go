package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//cmd := `df /myfiles | tail -n +2 | awk '{printf "{ \"part\": \"%s\", \"total\":  %s, \"used\": %s, \"free\": %s }", $6, $2, $3, $4}'`
	cmd := exec.Command("df", "/myfiles")
	stdout, _ := cmd.StdoutPipe()
	cmd2 := exec.Command("tail", "-n", "+2", stdout)
	fmt.Println(string(cmd))
}
