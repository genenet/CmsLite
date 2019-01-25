package controllers

import (
	"github.com/astaxie/beego"
)

type StandardController struct {
	beego.Controller
}

func (c *StandardController) Welcome() {
	c.TplName = "standard/welcome.tpl"
}
func (c *StandardController) Mainframe() {
	c.TplName = "standard/main.tpl"
}
func (c *StandardController) List() {
	c.TplName = "standard/list.tpl"
}
func (c *StandardController) Form() {
	c.TplName = "standard/form.tpl"
}
