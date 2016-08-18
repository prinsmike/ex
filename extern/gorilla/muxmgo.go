package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type Path struct {
	Path string "path"
}

type Templates struct {
	T []string
}

func main() {

	r := mux.NewRouter()

	paths := getPaths()
	for _, path := range paths {
		r.HandleFunc(path.Path, pathHandler)
	}

	fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
	http.Handle("/js/", fileServer)
	fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

func pathHandler(w http.ResponseWriter, req *http.Request) {

	templates := &Templates{[]string{
		`{{define "Items"}}<div><ul>{{range $v := .Paths}}<li><a href="http://localhost:8080{{$v.Path}}">{{$v.Path}}</a></li>{{end}}</ul></div>{{end}}`,
		`{{define "Content"}}<div>{{.Path}}</div><div>{{template "Items" .}}</div>{{end}}`,
		`<html><head><title>{{.Path}}</title></head><body>{{template "Content" .}}</body></html>`,
	}}

	var data = make(map[string]interface{})
	data["Path"] = req.URL.Path
	fmt.Println(data["Path"])
	paths := getPaths()
	data["Paths"] = paths
	t, err := ParseTemplates(templates.T...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	var data = make(map[string]string)
	data["scheme"] = req.URL.Scheme
	data["host"] = req.URL.Host
	data["path"] = req.URL.Path
	fmt.Println(w.Header())
	templates := &Templates{[]string{
		`{{define "Content"}}<h1>404</h1><div>{{.scheme}}://{{.host}}{{.path}} was not found</div>{{end}}`,
		`<html><head><title>{{.scheme}}://{{.host}}{{.path}}</title></head><body>{{template "Content" .}}</body></html>`,
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

func getPaths() []Path {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("mux")

	result := []Path{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	return result
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
