package routers

import (
	"light-apk/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/delete", &controllers.MainController{}, "get:Delete")
	beego.Router("/gl", &controllers.MainController{}, "get:GetLatest")
	beego.AutoRouter(&controllers.MainController{})
}
