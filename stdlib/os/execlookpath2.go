package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal("ls is not installed.")
	}
	fmt.Printf("ls is available at %s\n", path)
}
