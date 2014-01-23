// main.go
package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  // render html templates from templates directory
  m.Use(render.Renderer())

  m.Get("/t1", func(r render.Render) {
    r.HTML(200, "t1/t1", nil)
  })

  m.Get("/t2", func(r render.Render) {
  	r.HTML(200, "t2/t2", nil)
  })

  m.Run()
}