package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello-web/controllers:HomeController"] = append(beego.GlobalControllerRouter["hello-web/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Welcome",
            Router: `/home/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello-web/controllers:HomeController"] = append(beego.GlobalControllerRouter["hello-web/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Home",
            Router: `/home/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
