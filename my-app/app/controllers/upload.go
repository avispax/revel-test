package controllers

import (
	"reflect"

	"github.com/revel/revel"
)

type Upload struct {
	*revel.Controller
}

func (c *Upload) Image() revel.Result {
	aaa := len(c.Params.Files)
	println("omy:2:", aaa)

	image := c.Params.Files["fileInput"][0]
	println(reflect.TypeOf(image))
	println("omy:3:", image.Filename, ",", image.Size, ",", image.Header)

	redirect := c.Redirect("/")
	return redirect
}
