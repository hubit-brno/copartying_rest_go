package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  //"github.com/jinzhu/gorm"
  "./copartying_rest_go"
)
func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })
  m.Get("/parties", GetParties)
  m.Get("/party", GetParty)

  m.Run()
}


func GetParties(r render.Render) {
  var parties [1]party.Party
  parties[0] = party.Party{ID: 255, Name: "Párty", Description: "Velký popis Party"}
  r.JSON(200, parties)
}

func GetParty(r render.Render) {
  p := party.Party{ID: 255, Name: "Jedna párty", Description: "Velký popis Party"}
  r.JSON(200, p)
}
