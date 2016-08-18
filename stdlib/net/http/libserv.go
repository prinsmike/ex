package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./lib")))
	if err != nil {
		log.Printf("error running docs webserver: %v", err)
	}
}
