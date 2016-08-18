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
	session, e := mgo.Dial("localhost")
	if e != nil {
		panic(e)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("jstest").C("config")

	result := Person{}
	e = c.Find(nil).One(&result)
	if e != nil && e.Error() == "not found" {
		fmt.Println("The record could not be found.")
	} else if e != nil {
		panic(e)
	} else {
		fmt.Printf("%v\n", result)
	}
}
