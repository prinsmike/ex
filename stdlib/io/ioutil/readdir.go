package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	listing, err := ioutil.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range listing {
		fmt.Println(v.Name(), v.IsDir())
	}
}
