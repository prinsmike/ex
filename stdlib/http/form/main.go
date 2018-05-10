package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/submit-form", submitForm)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func showForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./form.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	if r.FormValue("checkme") == "" {
		fmt.Println("Check Me!!!")
	}
	for k, v := range r.Form {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ","))
	}
	fmt.Fprintf(w, "%#v", r.Form)
}
