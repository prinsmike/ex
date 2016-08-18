package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var counter int

func main() {
	fmt.Println(&counter)
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	counter = counter + 1
	fmt.Fprintf(res, fmt.Sprintf("%d", counter))
	fmt.Println(counter)
}
