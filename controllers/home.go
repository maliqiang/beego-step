package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct{
	beego.Controller
}

func (this *HomeController) URLMapping(){
	this.Mapping("Home",this.Home)
	this.Mapping("Welcome",this.Welcome)
}

// @router /home/:key [get]
func (c *HomeController) Home(){
	c.TplName = "home.tpl"
}


// @router /home/ [get]
func (c *HomeController) Welcome(){
	c.TplName = "home.tpl"
}