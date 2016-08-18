package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

func main() {
	fmt.Println(bson.NewObjectId())
}
