package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Name"] = "No file uploaded"
	c.Data["Size"] = "No file uploaded"
	c.TplName = "index.html"
}
