package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  //"github.com/martini-contrib/binding"
)

func NewCopartie(r render.Render){
  r.JSON(200, map[string]interface{}{"hello": "world"})
}

  func newCopartieItem(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  }

func GetCopartie(r render.Render){
  r.JSON(200, map[string]interface{}{"hello": "world"})
}

func UpdateCopartie(r render.Render) {
  r.JSON(200, map[string]interface{}{"hello": "world"})
}

func DeleteCopartie(r render.Render) {
  r.JSON(200, map[string]interface{}{"hello": "world"})
}
  



func main() {
  apiRoute := "/api/v1"

  m := martini.Classic()
  m.Use(render.Renderer())

  m.Group(apiRoute + "/coparties", func(r martini.Router) {
    r.Get("/:id", GetCopartie)
    r.Post("", NewCopartie)
    r.Post(":coparty_id/items/:id", newCopartieItem)
    r.Put("/:id", UpdateCopartie)
    r.Delete("/:id", DeleteCopartie)
  })

  m.Run()
}