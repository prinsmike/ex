package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("An error occured!")
	} else {
		fmt.Println(f)
	}
}
