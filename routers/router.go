package routers

import (
	"hello-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//注解路由
	beego.Include(&controllers.HomeController{})
}
