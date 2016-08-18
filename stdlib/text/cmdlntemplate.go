package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    string
}

func main() {
	sweaters := Inventory{os.Args[1], os.Args[2]}
	tmpl, err := template.New("test").Parse(os.Args[3])
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
