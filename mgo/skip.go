package main

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Person struct {
	Name    string
	Surname string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	result := []Person{}
	err = c.Find(nil).Sort("name", "surname").Limit(20).Skip(20).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(c.Count())
}
