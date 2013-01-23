package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/template", templateHandler)
	r.HandleFunc("/json", jsonHandler)
	r.HandleFunc("/{category}", categoryHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, Gopher!")
}

func templateHandler(res http.ResponseWriter, req *http.Request) {
	t := template.New("test")
	t, _ = t.Parse("Hello {{.Name}}!")
	t.Execute(res, map[string]string{"Name": "html/template"})
}

func categoryHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]
	t := template.New("test")
	t, _ = t.Parse("<h1>Hello {{.Category}}!</h1>")
	t.Execute(res, map[string]string{"Category": category})
}

func jsonHandler(res http.ResponseWriter, req *http.Request) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := string(b)
	fmt.Fprintf(res, s)
}
