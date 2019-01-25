package routers

import (
	"51hk_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/infos", &controllers.MainController{}, "*:Infos")
	beego.Router("/standard/welcome", &controllers.StandardController{}, "*:Welcome")
	beego.Router("/standard", &controllers.StandardController{}, "*:Mainframe")
	beego.Router("/standard/list", &controllers.StandardController{}, "*:List")
	beego.Router("/standard/form", &controllers.StandardController{}, "*:Form")
}
