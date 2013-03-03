package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

func main() {

	// startDate := "ISODate(\"2006-01-02T15:04\")"
	// endDate := "ISODate(\"2006-01-02T15:04\")"
	startDate, err := time.Parse("2006-01-02T15:04", "2013-01-01T12:00")
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("2006-01-02T15:04", "2013-02-19T14:30")
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("apache")

	var q = []bson.M{
		bson.M{"$match": bson.M{"reqTime": bson.M{"$gt": startDate, "$lt": endDate}}},
		bson.M{"$group": bson.M{"_id": "$serverName", "logCount": bson.M{"$sum": 1}, "bytesReceived": bson.M{"$sum": "$bytesReceived"}, "bytesSent": bson.M{"$sum": "$bytesSent"}}},
		bson.M{"$sort": bson.M{"logCount": 1}},
	}
	d := bson.D{{"aggregate", "log"}, {"pipeline", q}}

	var Stats = make(map[string]interface{})
	err = db.Run(d, &Stats)
	if err != nil {
		panic(err)
	}
	fmt.Println(Stats)
}
