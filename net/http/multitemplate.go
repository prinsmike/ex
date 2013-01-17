package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Templates struct {
	T []string
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Hello World"
	data["Body"] = "This is a Hello World example."
	data["Items"] = []string{"item 1", "item 2", "item 3"}

	templates := &Templates{[]string{
		`{{define "Items"}}<div><ul>{{range $v := .Items}}<li>{{$v}}</li>{{end}}</ul></div>{{end}}`,
		`{{define "Content"}}<div>{{.Body}}</div><div>{{template "Items" .}}</div>{{end}}`,
		`<html><head><title>{{.Title}}</title></head><body>{{template "Content" .}}</body></html>`,
	}}

	t, err := ParseTemplates(templates.T...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func ParseTemplates(templates ...string) (*template.Template, error) {
	return parseTemplates(nil, templates...)
}

func parseTemplates(t *template.Template, templates ...string) (*template.Template, error) {
	if len(templates) == 0 {
		return nil, fmt.Errorf("No templates specified in call to ParseTemplates.")
	}
	name := "temporary.template.name"
	for _, tv := range templates {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		tmpl = t

		_, err := tmpl.Parse(tv)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
