package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Charset: "UTF8",
	}))

	m.Get("/info", func(r render.Render, params martini.Params) { r.JSON(200, Version{12}) })

	m.Run()
}
