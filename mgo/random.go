package main

import (
	"fmt"
	"github.com/prinsmike/passgo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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

	k := []byte("bcdfghjklmnpqrstvwxyz")
	v := []byte("aeiou")

	gen := passgo.NewGenerator(k, v, nil, nil, false, 0)

	for i := 0; i <= 10; i++ {
		name, _ := gen.NewPassword(5, 0, 0)
		surname, _ := gen.NewPassword(8, 0, 0)
		err = c.Insert(&Person{name, surname})
		if err != nil {
			panic(err)
		}
	}

	result := []Person{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(c.Count())

}
