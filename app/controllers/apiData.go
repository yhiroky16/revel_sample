package controllers

import (
	"log"
	"github.com/revel/revel"
)

type ApiDataController struct {
	*revel.Controller
}

type ApiData struct {
	ApiDataController
}

func (c ApiData) Show(id int) revel.Result {
	log.Printf("Show method call! ¥n")
	return c.Render()
}
