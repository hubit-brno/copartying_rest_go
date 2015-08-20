package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"
  "./copartying_rest_go"
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
    r.Get("/:id&secret_token=:admin_secret_token", GetCopartie)
    r.Post("", NewCopartie)
    r.Post(":coparty_id/items/:id&secret_token=:user_secret_token", newCopartieItem)
    r.Put("/:id&secret_token=:admin_secret_token", UpdateCopartie)
    r.Delete("/:id&secret_token=:admin_secret_token", DeleteCopartie)
  })

  m.Run()
}