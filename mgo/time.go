package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

func main() {

	id := bson.NewObjectId()
	fmt.Println(id.Time())
}
