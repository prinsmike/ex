package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	d := make(map[string]interface{})
	d["Title"] = "Hello"
	d["Items"] = []string{"item 1", "item 2"}
	t := template.New("Template")
	_, err := t.Parse("<html><head><title>{{.Title}}</title></head><body><h1>{{.Title}}</h1><ul>{{range $v := .Items}}<li>{{$v}}</li>{{end}}</ul></body></html>")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
