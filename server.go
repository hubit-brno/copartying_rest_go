package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "log"
)
import ara "github.com/diegogub/aranGO"
import cop "hub/copartying"

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  s, err := ara.Connect("http://127.0.0.1:8529", "hub2", "IFeelGood2", true)
  if err != nil {
    panic(err)
  }

  cop.Init(s, "hub2")

  // var party cop.PartyDoc
  // party.Name = "Woot Cop"
  // key := cop.CreateParty(party)
  // log.Println("Key: " + key)

  key := "471727256"
  var party2 cop.PartyDoc = cop.GetParty(key)
  log.Println("Key: " + key)
  log.Println("Retrieved name: " + party2.Name)


  // m.Get("/", func(r render.Render) {
  //   r.JSON(200, map[string]interface{}{"hello": "world"})
  // })


  // m.Run()
}
