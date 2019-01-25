package controllers

import (
	da "51hk_beego/dao"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var ipageCount int
	infos, _ := da.GetAll(-1, -1, "", 1, 8, &ipageCount)
	c.Data["result"] = infos.Data

	c.TplName = "home/index.tpl"
}
func (c *MainController) Infos() {

	var ipageCount int
	infos, _ := da.GetAll(-1, -1, "", 1, 8, &ipageCount)
	c.Data["json"] = infos.Data
	c.ServeJSON()
}
