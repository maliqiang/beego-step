package main

import (
	_ "hello-web/routers"
	"github.com/astaxie/beego"
)

func main() {
	//添加静态目录访问
	beego.SetStaticPath("/download", "download")
	beego.Run()
}