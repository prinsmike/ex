package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", listDBs).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func listDBs(w http.ResponseWriter, req *http.Request) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	databases := make(map[string][]string)
	databases["list"], err = session.DatabaseNames()
	if err != nil {
		log.Fatalln(err)
	}
	j, err := json.Marshal(databases)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, string(j))
	fmt.Println(string(j))
}
