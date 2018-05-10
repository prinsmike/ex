package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var envMap map[string]string
	var err error
	envMap, err = godotenv.Read("substitution.sh")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", envMap)
}
