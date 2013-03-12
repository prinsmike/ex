package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", listDBs).Methods("GET")
	r.HandleFunc("/{dbname}", listCollections).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func listDBs(w http.ResponseWriter, req *http.Request) {
	mongo := new(Mongo)
	mongo.URL = "localhost"
	mongo.Session = mongo.GetSession()

	dbmap := make(map[string][]string)
	var err error
	dbmap["databases"], err = mongo.Session.DatabaseNames()
	if err != nil {
		log.Fatalln(err)
	}
	j, err := json.Marshal(dbmap)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, string(j))
	fmt.Println(string(j))
}

func listCollections(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	colmap := make(map[string][]string)

	var err error
	_, db := NewMongoConnection("localhost", vars["dbname"])
	colmap["collections"], err = db.CollectionNames()
	if err != nil {
		log.Fatalln(err)
	}
	j, err := json.Marshal(colmap)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, string(j))
	fmt.Println(string(j))
}

type Mongo struct {
	URL      string
	DbStr    string
	Session  *mgo.Session
	Database *mgo.Database
}

// A convenient alias for bson.M so we don't have to import bson everywhere.
type M bson.M

func (mongo *Mongo) GetSession() *mgo.Session {
	if mongo.Session == nil {
		var err error
		mongo.Session, err = mgo.Dial(mongo.URL)
		if err != nil {
			panic(err)
		}
	}
	return mongo.Session.Clone()
}

func (mongo *Mongo) GetDb() *mgo.Database {
	return mongo.Session.DB(mongo.DbStr)
}

func NewMongoConnection(url, db string) (session *mgo.Session, database *mgo.Database) {
	mongo := new(Mongo)
	mongo.URL = url
	mongo.DbStr = db
	session = mongo.GetSession()
	database = mongo.GetDb()
	return
}
