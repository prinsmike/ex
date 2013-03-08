package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", pathHandler)
	r.HandleFunc("/{path:.*}", pathHandler).Host("test.local")
	r.HandleFunc("/{path:.*}", pathHandler).Host("test1.local")
	r.HandleFunc("/{path:.*}", noHostHandler).Host("{host:.*}")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func pathHandler(w http.ResponseWriter, req *http.Request) {
	var data = make(map[string]string)
	data["method"] = req.Method
	data["proto"] = req.Proto
	data["requri"] = req.RequestURI
	data["host"] = req.Host
	data["path"] = req.URL.Path
	data["useragent"] = req.UserAgent()
	templates := &Templates{[]string{
		`{{define "Content"}}<div>Method: {{.method}}</div><div>{{.proto}}</div><div>Host: {{.host}}</div><div>Path: {{.path}}</div><div>Request URI: {{.requri}}</div><div>UserAgent: {{.useragent}}</div>{{end}}`,
		`<html><head><title>{{.host}}</title></head><body>{{template "Content" .}}</body></html>`,
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

func noHostHandler(w http.ResponseWriter, req *http.Request) {
	var data = make(map[string]string)
	data["method"] = req.Method
	data["proto"] = req.Proto
	data["requri"] = req.RequestURI
	data["host"] = req.Host
	data["path"] = req.URL.Path
	data["useragent"] = req.UserAgent()
	templates := &Templates{[]string{
		`{{define "Content"}}<div>Method: {{.method}}</div><div>{{.proto}}</div><div>Host: {{.host}}</div><div>Path: {{.path}}</div><div>Request URI: {{.requri}}</div><div>UserAgent: {{.useragent}}</div>{{end}}`,
		`<html><head><title>{{.host}}</title></head><body><h1>No site on host: {{.host}}</h1><div>{{template "Content" .}}</div></body></html>`,
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

type Templates struct {
	T []string
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
