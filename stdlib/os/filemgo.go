package main

import (
	"bufio"
	"encoding/json"
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Id            bson.ObjectId "_id,omitempty"
	IP            string        "ip"
	Time          int           "time"
	ReqProtocol   string        "reqProtocol"
	ReqMethod     string        "reqMethod"
	QueryString   string        "queryString"
	LastRequest   int           "lastRequest"
	ReqTime       time.Time     "reqTime"
	Path          string        "path"
	ServerName    string        "serverName"
	BytesReceived int           "bytesReceived"
	BytesSent     int           "bytesSent"
	Referer       string        "Referer"
	UserAgent     string        "userAgent"
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("apache").C("log")

	f, err := os.Open("apache.log")
	if err != nil {
		panic(err)
	}
	data := bufio.NewReader(f)
	for {
		id := bson.NewObjectId()
		l := new(LogEntry)
		l.Id = id
		line, err := data.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for i, ch := range line {
			if ch == 92 {
				line[i] = 124
			}
		}
		err = json.Unmarshal(line, &l)
		if err != nil {
			log.Println(err)
		}
		err = c.Insert(&l)
		if err != nil {
			log.Println(err)
		}
	}
}
