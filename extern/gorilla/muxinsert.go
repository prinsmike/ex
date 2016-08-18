package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

type Path struct {
	Id   bson.ObjectId "_id,omitempty"
	Path string        "path"
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You must provide a path!")
		os.Exit(1)
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("mux")

	index := mgo.Index{
		Key:        []string{"path"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		log.Println(err)
	}

	id := bson.NewObjectId()
	p := new(Path)
	p.Id = id
	p.Path = os.Args[1]
	err = c.Insert(&p)
	if err != nil {
		panic(err)
	}
}
