package main

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

type Person struct {
	Id        bson.ObjectId "_id,omitempty"
	FirstName string        "firstName"
	LastName  string        "lastName"
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("stdinmgo")

	dec := json.NewDecoder(os.Stdin)

	for {
		id := bson.NewObjectId()
		p := new(Person)
		p.Id = id
		if err := dec.Decode(&p); err != nil {
			log.Println(err)
			return
		}
		err = c.Insert(&p)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted document.")
	}
}
