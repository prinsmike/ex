package main

import (
	"encoding/json"
	"fmt"
	"github.com/DisposaBoy/JsonConfigReader"
	"os"
)

func main() {
	var v interface{}
	f, _ := os.Open("settings.json")
	// wrap our reader before passing it to the json decoder
	r := JsonConfigReader.New(f)
	json.NewDecoder(r).Decode(&v)
	fmt.Println(v)
}
