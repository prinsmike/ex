package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Template struct {
	Name string
	T    string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	templates := []interface{}{
		&Template{`List`, `{{define "Items"}}<div><ul>{{range $v := .Items}}<li>{{$v}}</li>{{end}}</ul></div>{{end}}`},
		&Template{`Content`, `{{define "Content"}}<div>{{.Body}}</div><div>{{template "Items" .}}</div>{{end}}`},
		&Template{`Base`, `<html><head><title>{{.Title}}</title></head><body>{{template "Content" .}}</body></html>`},
	}

	c := session.DB("test").C("templates")
	err = c.Insert(templates...)
	if err != nil {
		panic(err)
	}

	result := Template{}
	err = c.Find(bson.M{"name": "Base"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Template:", result.T)
}
