package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	hello := os.Getenv("HELLO")
	bye := os.Getenv("BYE")
	fmt.Println(hello)
	fmt.Println(bye)
}
