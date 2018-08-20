package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zd112_api/controllers:AppController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:AppController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Advert",
			Router: `/advert`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:AppController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Stop",
			Router: `/stop`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:AppController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:BlockChainController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:BlockChainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:BlockChainController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:BlockChainController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:HomeController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:HomeController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zd112_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["zd112_api/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Multipart",
			Router: `/multipart`,
			AllowHTTPMethods: []string{"post"},
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
