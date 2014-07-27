package main

import (
	"fmt"
	"net/http"
	"os"
)

type Path struct {
	Host, Path string
}

var pages = make(map[Path]string)

func httpHandler(w http.ResponseWriter, req *http.Request) {
	if content, ok := pages[Path{req.Host, req.URL.Path}]; ok {
		fmt.Fprintf(w, content)
	} else {
		fmt.Fprintf(w, "Could not find the requested URL")
	}
}

func main() {

	pages[Path{"localhost:8080", "/"}] = "<h1>Welcome</h1>"
	pages[Path{"localhost:8080", "/test"}] = "<h1>Test</h1>"
	pages[Path{"test.locl", "/test"}]

	http.HandleFunc("/", httpHandler)

	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println("Could not start web server!")
	}
}
