// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zd112_api/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			), ),
		beego.NSNamespace("/blockchain",
			beego.NSInclude(
				&controllers.BlockChainController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.ErrorController(&controllers.ErrorController{})
}
