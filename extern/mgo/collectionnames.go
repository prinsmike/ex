package main

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	names, err := session.DB("test").CollectionNames()
	if err != nil {
		panic(err)
	}

	fmt.Println(names)
}
