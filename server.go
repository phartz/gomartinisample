package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func getInfoHandler(w http.ResponseWriter, params martini.Params) {
	js, err := json.Marshal(Version{4711})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Charset: "UTF8",
	}))

	m.Get("/renderer", func(r render.Render, params martini.Params) { r.JSON(200, Version{12}) })
	m.Get("/info", getInfoHandler)

	m.Run()
}
