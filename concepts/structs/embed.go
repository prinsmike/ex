package main

import "fmt"

type Foo struct {
	name string
	opt  Options
}

type Options struct {
	Max              int64
	MaxTaskNameWidth int
	AlwaysUpdate     bool
	// etc  
}

func NewFoo(name string, opt Options) *Foo {
	return &Foo{name, opt}
}

func main() {

	options := Options{123, 34, true}

	foo := NewFoo("Hello", options)

	fmt.Println("Foo >>", foo)
}
