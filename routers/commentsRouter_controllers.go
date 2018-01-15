package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zd112_api/controllers:HomeController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UserController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
