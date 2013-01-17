package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", servFunc)
	fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
	http.Handle("/js/", fileServer)
	fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)
	fileServer = http.StripPrefix("/lib/", http.FileServer(http.Dir("lib")))
	http.Handle("/lib/", fileServer)

	err := http.ListenAndServe(":8080", nil)
	checkError(err)
}

func servFunc(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
