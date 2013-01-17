package main

import (
	"html/template"
	"net/http"
)

type Content struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	content := &Content{"Hello Title", "Hello body."}
	tmpl, err := template.New("test").Parse("<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
}
