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

func main() {

	pages[Path{"localhost:8080", "/"}] = "<h1>Welcome</h1>"
	pages[Path{"localhost:8080", "/test"}] = "<h1>Test</h1>"

	http.HandleFunc("/", servFunc)

	err := http.ListenAndServe(":8080", nil)
	checkError(err)
}

func servFunc(w http.ResponseWriter, req *http.Request) {
	host := req.Host
	path := req.URL.Path
	if content, ok := pages[Path{host, path}]; ok {
		fmt.Fprintf(w, content)
	} else {
		fmt.Fprintf(w, "Could not find the requested URL")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
