package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	var m Message
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("An error occured!")
	} else {
		fmt.Println(m)
	}
}
