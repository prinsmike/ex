package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./nestedvars.sh")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getEnv("V4"))
}

func getEnv(v string) string {
	var s string
	s = os.Getenv(v)
	log.Println(s)
	if strings.HasPrefix(s, "$") {
		s = getEnv(strings.TrimPrefix(s, "$"))
	}
	return s
}
