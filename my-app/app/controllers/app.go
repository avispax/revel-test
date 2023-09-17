package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	testString := "testString"
	return c.Render(testString)
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}
