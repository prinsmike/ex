package main

import (
	"net/http"
	"github.com/codegangsta/martini"
	"fmt"
	"log"
)

func main() {
	m := martini.Classic()
	//fmt.Printf("%v\n", m)
	m.Get(".*", func(req *http.Request) string {
		fmt.Printf("%#v\n", req)
		fmt.Println()
		fmt.Printf("%#v\n", req.URL)
		return "Hello World!"
	})
	
	//fmt.Printf("%#v\n", tmp)

	log.Fatal(http.ListenAndServe(":6161", m))
}