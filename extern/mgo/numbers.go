package main

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Number struct {
	N int
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("numbers")

	index := mgo.Index{
		Key:        []string{"n"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	_ = c.EnsureIndex(index)

	count, err := c.Count()
	if err != nil {
		panic(err)
	}
	count++

	for i := count; i <= (count + 19); i++ {
		err = c.Insert(&Number{i})
		if err != nil {
			panic(err)
		}
	}

	result := []Number{}
	err = c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(c.Count())

}
