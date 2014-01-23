package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/form"
)

type BlogPost struct {
	Title   string `form:"title,required"`
	Content string `form:"content"`
}

func main() {
	m := martini.Classic()

	m.Post("/blog", form.Form(&BlogPost{}), func(blogpost *BlogPost) string {
		return blogpost.Title
	})

	m.Run()
}
